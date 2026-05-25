--[[
  JobOps latex-pdf — OMERS Lua filter (Tier 2 parity).

  Maps the JobOps markdown authoring contract onto the hand-built OMERS LaTeX
  vocabulary so resumes and cover letters render to the gold standard.

  Activated by the skill via:  pandoc ... --lua-filter=omers-filter.lua -M doctype=<doctype>

  Behaviour by doctype:
    resume       first H1 (name + post-noms) + following tagline/contact lines
                 -> centered header block; H2 -> \section; H3 + italic subline
                 -> \role{COMPANY}{dates} + \subrole{title — location};
                 markdown tables -> navy-header + zebra longtable.
    coverletter  first H1 + contact line -> left letterhead (name + navy rule +
                 contact); H2 -> \section; markdown tables -> navy-header + zebra
                 longtable (the Requirements Alignment table).
    document     PASS-THROUGH. The filter does nothing; pandoc's default heading
                 mapping + the document preamble handle styling.

  Tolerant by design: if an H3 subline has no " | " dates, the whole line
  becomes the subrole with no right-aligned dates; if no header material follows
  the H1, only the name renders.
]]

local stringify = pandoc.utils.stringify

-- Render a list of inlines to a LaTeX string (preserves links, emphasis, etc.).
local function inlines_latex(inlines)
  local s = pandoc.write(pandoc.Pandoc({ pandoc.Plain(inlines) }), 'latex')
  return (s:gsub('%s+$', ''))
end

-- Render a list of blocks (e.g. a table cell's contents) to a LaTeX string.
local function blocks_latex(blocks)
  local s = pandoc.write(pandoc.Pandoc(blocks), 'latex')
  return (s:gsub('%s+$', ''))
end

-- Escape the handful of LaTeX specials that show up in plain-text fields
-- (names, dates, sublines). Unicode passes through untouched under xelatex.
local function tex_escape(s)
  s = s:gsub('\\', '\\textbackslash{}')
  s = s:gsub('([%%%$%#%&%_%{%}])', '\\%1')
  s = s:gsub('~', '\\textasciitilde{}')
  s = s:gsub('%^', '\\textasciicircum{}')
  return s
end

local function split(str, sep)
  local out, pat = {}, '(.-)' .. sep
  local last_end = 1
  local s, e, cap = str:find(pat, 1)
  while s do
    table.insert(out, cap)
    last_end = e + 1
    s, e, cap = str:find(pat, last_end)
  end
  table.insert(out, str:sub(last_end))
  return out
end

local function trim(s) return (s:gsub('^%s+', ''):gsub('%s+$', '')) end

-- Does a Para/Plain look like part of the header block (tagline or contact)?
local function is_header_material(block)
  if block.t ~= 'Para' and block.t ~= 'Plain' then return false end
  local has_link = false
  pandoc.walk_block(block, { Link = function() has_link = true; return nil end })
  if has_link then return true end
  local txt = stringify(block)
  if txt:find('@') or txt:find('https?://') or txt:lower():find('linkedin') then return true end
  if txt:find('•') or txt:find('|') then return true end
  if #block.content == 1 and block.content[1].t == 'Strong' then return true end
  return false
end

local function has_link_or_at(block)
  local found = false
  pandoc.walk_block(block, { Link = function() found = true; return nil end })
  if found then return true end
  return stringify(block):find('@') ~= nil
end

-- First inline of a Para is Emph => treat as an italic subline.
local function is_italic_para(block)
  if block.t ~= 'Para' and block.t ~= 'Plain' then return false end
  for _, il in ipairs(block.content) do
    if il.t == 'Space' or il.t == 'SoftBreak' then
      -- skip leading spaces
    else
      return il.t == 'Emph'
    end
  end
  return false
end

-- Split an H1 "Name, CFA, FRICS" into {name, postnoms}.
local function split_name(full)
  local comma = full:find(',')
  if comma then
    return trim(full:sub(1, comma - 1)), trim(full:sub(comma)) -- postnoms keeps leading comma
  end
  return trim(full), ''
end

local function build_header(doctype, name_text, consumed)
  local name, postnoms = split_name(name_text)
  -- classify consumed paras into tagline (no link) vs contact (link/@)
  local tagline_latex, contact_latex = nil, nil
  for _, blk in ipairs(consumed) do
    if has_link_or_at(blk) then
      contact_latex = inlines_latex(blk.content)
    else
      tagline_latex = inlines_latex(blk.content)
    end
  end

  if doctype == 'resume' then
    local lines = {}
    table.insert(lines, '\\begin{center}')
    local namepart = '{\\headfont\\bfseries\\color{accent}\\fontsize{24}{27}\\selectfont\\addfontfeature{LetterSpace=1.5}'
      .. tex_escape(name) .. '}'
    if postnoms ~= '' then
      namepart = namepart .. '{\\headfont\\color{muted}\\fontsize{14}{17}\\selectfont '
        .. tex_escape(postnoms) .. '}'
    end
    table.insert(lines, namepart .. '\\\\[3pt]')
    if tagline_latex then
      table.insert(lines, '{\\headfont\\color{muted}\\fontsize{11}{13}\\selectfont\\addfontfeature{LetterSpace=2.0}'
        .. tagline_latex .. '}\\\\[5pt]')
    end
    if contact_latex then
      table.insert(lines, '{\\small ' .. contact_latex .. '}')
    end
    table.insert(lines, '\\end{center}')
    table.insert(lines, '\\vspace{2pt}')
    return table.concat(lines, '\n')
  else -- coverletter letterhead
    local lines = {}
    local namepart = '{\\headfont\\bfseries\\color{accent}\\fontsize{25}{28}\\selectfont\\addfontfeature{LetterSpace=1.5}'
      .. tex_escape(name) .. '}'
    if postnoms ~= '' then
      namepart = namepart .. '{\\headfont\\color{muted}\\fontsize{13}{16}\\selectfont '
        .. tex_escape(postnoms) .. '}'
    end
    table.insert(lines, namepart .. '\\\\[5pt]')
    table.insert(lines, '{\\color{accent}\\rule{\\linewidth}{1pt}}\\\\[3pt]')
    if contact_latex then
      table.insert(lines, '{\\headfont\\color{muted}\\small ' .. contact_latex .. '}')
    end
    table.insert(lines, '\\vspace{1.2em}')
    return table.concat(lines, '\n')
  end
end

-- Build an OMERS-styled longtable (navy header row + zebra body) from a Table.
local function build_table(tbl, doctype)
  local ncol = #tbl.colspecs
  if ncol == 0 then return nil end

  -- Column fractions. The 2-col cover-letter requirements table is a known
  -- shape, so force the OMERS 29/71 split for parity. Otherwise honour
  -- pandoc-provided relative widths, falling back to an equal split.
  local fracs = {}
  if ncol == 2 and doctype == 'coverletter' then
    fracs = { 0.29, 0.71 }
  else
    local have_widths = false
    for _, cs in ipairs(tbl.colspecs) do
      local w = cs[2]
      if type(w) == 'number' and w > 0 then have_widths = true end
    end
    for _, cs in ipairs(tbl.colspecs) do
      fracs[#fracs + 1] = (have_widths and cs[2]) or (1.0 / ncol)
    end
  end

  local colspec = {}
  for k = 1, ncol do
    -- leave a little slack (×0.96) so inter-column tabcolsep never overflows
    colspec[k] = string.format(
      '>{\\raggedright\\arraybackslash}p{(\\columnwidth - 2\\tabcolsep)*\\real{%.4f}}',
      fracs[k] * 0.96)
  end

  local function row_cells(row)
    local cells = {}
    for _, cell in ipairs(row.cells) do
      cells[#cells + 1] = blocks_latex(cell.contents)
    end
    return cells
  end

  local out = {}
  table.insert(out, '{\\setlength{\\LTpre}{2pt}\\setlength{\\LTpost}{0pt}')
  table.insert(out, '\\begin{longtable}{' .. table.concat(colspec, ' ') .. '}')

  -- header row(s)
  for _, row in ipairs(tbl.head.rows) do
    local cells = row_cells(row)
    local hcells = {}
    for _, c in ipairs(cells) do hcells[#hcells + 1] = '\\thcell{' .. c .. '}' end
    table.insert(out, '\\rowcolor{accent}' .. table.concat(hcells, ' & ') .. '\\\\')
  end
  table.insert(out, '\\endhead')

  -- body rows with zebra striping
  local r = 0
  for _, body in ipairs(tbl.bodies) do
    for _, row in ipairs(body.body) do
      r = r + 1
      local cells = row_cells(row)
      local prefix = (r % 2 == 0) and '\\rowshade ' or ''
      table.insert(out, prefix .. table.concat(cells, ' & ') .. '\\\\')
    end
  end

  table.insert(out, '\\end{longtable}}')
  return table.concat(out, '\n')
end

function Pandoc(doc)
  local doctype = stringify(doc.meta.doctype or '')
  if doctype ~= 'resume' and doctype ~= 'coverletter' then
    return nil -- pass-through for `document` and anything else
  end

  local blocks = doc.blocks
  local out = {}
  local i = 1
  local header_done = false

  while i <= #blocks do
    local b = blocks[i]

    if b.t == 'Header' and b.level == 1 and not header_done then
      local name_text = stringify(b.content)
      i = i + 1
      local consumed = {}
      while i <= #blocks and #consumed < 2 and is_header_material(blocks[i]) do
        consumed[#consumed + 1] = blocks[i]
        i = i + 1
      end
      out[#out + 1] = pandoc.RawBlock('latex', build_header(doctype, name_text, consumed))
      header_done = true

    elseif b.t == 'Header' and b.level == 3 and doctype == 'resume' then
      local company = tex_escape(stringify(b.content))
      local dates, subrole = '', ''
      if blocks[i + 1] and is_italic_para(blocks[i + 1]) then
        local sub = trim(stringify(blocks[i + 1]))
        local parts = split(sub, ' | ')
        if #parts >= 2 then
          dates = trim(parts[#parts])
          table.remove(parts, #parts)
          for k = 1, #parts do parts[k] = trim(parts[k]) end
          subrole = table.concat(parts, ' — ')
        else
          subrole = sub
        end
        i = i + 1 -- consume the subline
      end
      out[#out + 1] = pandoc.RawBlock('latex', '\\role{' .. company .. '}{' .. tex_escape(dates) .. '}')
      if subrole ~= '' then
        out[#out + 1] = pandoc.RawBlock('latex', '\\subrole{' .. tex_escape(subrole) .. '}')
      end
      i = i + 1

    elseif b.t == 'Header' then
      out[#out + 1] = pandoc.RawBlock('latex', '\\section{' .. inlines_latex(b.content):gsub('\n', ' ') .. '}')
      i = i + 1

    elseif b.t == 'Table' then
      local t = build_table(b, doctype)
      out[#out + 1] = t and pandoc.RawBlock('latex', t) or b
      i = i + 1

    else
      out[#out + 1] = b
      i = i + 1
    end
  end

  doc.blocks = out
  return doc
end
