# conways_game_of_life_tui
conways game of life as a tui using golang

## Usage
if you have go installed, you can run the following command to run the program (assuming you have the repo cloned and are in the root directory)
```bash
go run main.go
```

if you have just the binary, you can run the following command to run the program
```bash
./conways_game_of_life_tui
```

to check the help menu, you can run the following command
```bash
./conways_game_of_life_tui -h
```

## Controls
Hitting ESC at any time will bring up the menu. You can use the arrow keys to navigate the menu and hit ENTER to select an option.

When in the config menu you can use TAB and Shift+TAB to navigate between the different fields. You can type in fields to change the number. Pressing enter on a checkbox or button will activate it.

If you want to use a custom initial board you go to the config page, check the Custom Board checkbox, and then use the Custom Board button to go to a grid table to set up the custom board. Pressing TAB while in the custom board grid will save what you filled in.

When you are done editing the config hit ENTER on the Done button to start the game.

