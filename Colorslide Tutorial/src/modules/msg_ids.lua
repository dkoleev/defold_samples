local M = {}

-- Engine actions
M.INIT = hash("init")
M.FINAL = hash("final")
M.ENABLE = hash("enable")
M.DISABLE = hash("disable")
M.UNLOAD = hash("unload")
M.ACQUIRE_INPUT_FOCUS = hash("acquire_input_focus")
M.TOUCH = hash("touch")

-- Game actions
M.LOAD = hash("load")
M.LOAD_LEVEL = hash("load_level")
M.UNLOAD_LEVEL = hash("unload_level")
M.NEXT_LEVEL = hash("next_level")
M.PROXY_LOADED = hash("proxy_loaded")
M.SHOW_LEVEL_SELECT = hash("show_level_select")
M.HIDE_LEVEL_SELECT = hash("hide_level_select")
M.LEVEL_COMPLETED = hash("level_completed")
M.SHOW_START = hash("show_start")
M.HIDE_START = hash("hide_start")

return M