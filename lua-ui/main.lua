-- main.lua

local buttons = {}

-- button factory
local function newButton(text, x, y, w, h, onClick)
	return {
		text = text,
		x = x,
		y = y,
		w = w,
		h = h,
		hovered = false,
		onClick = onClick,
	}
end

local function isHovered(btn)
	local mx, my = love.mouse.getPosition()
	return mx >= btn.x and mx <= btn.x + btn.w and my >= btn.y and my <= btn.y + btn.h
end

local function drawButton(btn)
	if btn.hovered then
		love.graphics.setColor(0.3, 0.6, 1) -- blue when hovered
	els
		love.graphics.setColor(0.2, 0.2, 0.2) -- dark when idle
	end

	love.graphics.rectangle("fill", btn.x, btn.y, btn.w, btn.h, 8, 8) -- 8px rounded corners

	love.graphics.setColor(1, 1, 1)
	love.graphics.printf(btn.text, btn.x, btn.y + btn.h / 2 - 8, btn.w, "center")
end

-- screens
local currentScreen = "menu"

local function goTo(screen)
	currentScreen = screen
end

-- love callbacks
function love.load()
	love.window.setTitle("Chess")
	love.window.setMode(800, 600)

	local cx = 800 / 2 - 150 -- center x
	buttons.play = newButton("Play", cx, 200, 300, 50, function()
		goTo("game")
	end)
	buttons.settings = newButton("Settings", cx, 270, 300, 50, function()
		goTo("settings")
	end)
	buttons.quit = newButton("Quit", cx, 340, 300, 50, function()
		love.event.quit()
	end)
end

function love.update(dt)
	for _, btn in pairs(buttons) do
		btn.hovered = isHovered(btn)
	end
end

function love.draw()
	if currentScreen == "menu" then
		-- background
		love.graphics.setColor(0.1, 0.1, 0.1)
		love.graphics.rectangle("fill", 0, 0, 800, 600)

		-- title
		love.graphics.setColor(1, 1, 1)
		love.graphics.printf("CHESS", 0, 100, 800, "center")

		-- buttons
		for _, btn in pairs(buttons) do
			drawButton(btn)
		end
	elseif currentScreen == "game" then
		love.graphics.setColor(1, 1, 1)
		love.graphics.printf("Game screen - press ESC to go back", 0, 280, 800, "center")
	elseif currentScreen == "settings" then
		love.graphics.setColor(1, 1, 1)
		love.graphics.printf("Settings screen - press ESC to go back", 0, 280, 800, "center")
	end
end

function love.mousepressed(x, y, button)
	if button == 1 then -- left click
		for _, btn in pairs(buttons) do
			if isHovered(btn) then
				btn.onClick()
			end
		end
	end
end

function love.keypressed(key)
	if key == "escape" then
		goTo("menu")
	end
end
