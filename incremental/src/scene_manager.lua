--- Scene management module for Defold
-- Handles loading/unloading collection proxies and level transitions

local log = require("log.log")
local logger = log.get_logger("scene_manager")

local M = {
    current_level = nil,
    loaded_scenes = {},
    levels_switch_transition_delay = 0
}

--- Find the index of a scene in loaded_scenes table
---@param scene string Scene path to find
---@return number|nil index Index if found, nil otherwise
local function find_scene_index(scene)
    for i, loaded_scene in ipairs(M.loaded_scenes) do
        if loaded_scene == scene then
            return i
        end
    end

    return nil
end

--- Load a scene (collection proxy)
---@param scene string Scene path (e.g., "/levels/level1#proxy")
---@return boolean success True if loaded, false if already loaded
function M.load_scene(scene)
    local scene_index = find_scene_index(scene)
    if scene_index then
        logger.error("Scene" .. tostring(scene) .. " is already loaded")

        return false
    else
        table.insert(M.loaded_scenes, scene)
        msg.post(scene, "load")


        logger:debug("Load scene: " .. scene)

        return true
    end
end

--- Unload a scene (collection proxy)
---@param scene string Scene path to unload
---@return boolean success True if unloaded, false if not loaded
function M.unload_scene(scene)
    local scene_index = find_scene_index(scene)

    if not scene_index then
        logger.error("Can't unload scene: " .. tostring(scene) .. " It is not loaded")

        return false
    end

    table.remove(M.loaded_scenes, scene_index)
    msg.post(scene, "unload")
    
    logger:debug("Unload scene: " .. scene)

    return true
end

--- Load a level and set it as current
---@param level string Level scene path
function M.load_level(level)
    M.current_level = level
    M.load_scene(level)
end

--- Unload the current level
function M.unload_level()
    if M.current_level then
        M.unload_scene(M.current_level)
        M.current_level = nil
    end
end

--- Switch to a new level immediately (unload then load)
---@param new_level string Scene path to switch to
function M.switch_level(new_level)
    M.unload_level()
    M.load_level(new_level)
end

--- Switch to a new level with optional transition delay
---@param new_level string Scene path to load (e.g., "/levels/level2#proxy")
---@param on_complete? fun() Callback when transition completes
---@param delay? number Delay in seconds before unloading old level
function M.switch_level_async(new_level, on_complete, delay)
    local old_level = M.current_level
    local wait_time = delay or M.levels_switch_transition_delay

    M.load_level(new_level)

    if old_level then
        timer.delay(wait_time, false, function()
            M.unload_scene(old_level)
            if on_complete then
                on_complete()
            end
        end)
    elseif on_complete then
        on_complete()
    end
end

return M
