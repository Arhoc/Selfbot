import discum
import os

bot = discum.Client(token=os.getenv("TOKEN"), log=False)

@bot.gateway.command
def command(resp):
    if resp.event.ready_supplemental:
        user = bot.gateway.session.user
        print(f"Logged in as {user['name']}{user['discriminator']}")
        os.system("neofetch")

    if resp.event.message:
        m = resp.parsed.auto()
        guildID = m['guild_id'] if 'guild_id' in m else None #because DMs are technically channels too
        channelID = m['channel_id']
        username = m['author']['username']
        discriminator = m['author']['discriminator']
        content = m['content']

        print("> guild {} channel {} | {}#{}: {}".format(guildID, channelID, username, discriminator, content))

bot.gateway.run(auto_reconnect=True)