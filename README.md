# Gimme Wisdom

Golang telegram wisdoms bot. Wisdoms are getting parsed from wisdoms.txt file, splitted by **". -"** pattern. To add a new wisdom, please add two line breaks ("\n") after the last quote block, then put a new quote and author name divided by **". -"** 

## Enviromental variables

**PUBLIC_URL**

**PORT** 

**TOKEN**

## Deploy on Heroku

1. Create a new Telegram bot and obtain API TOKEN from BotFather.
2. Add new application APP_NAME
3. Set enviromental variables pointed above
4. Run _heroku container:push web -a APP_NAME_
5. Run _heroku container:release web -a APP_NAME_