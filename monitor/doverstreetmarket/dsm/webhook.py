import os
from datetime import datetime
from discord_webhook import DiscordWebhook, DiscordEmbed


def webhook_time_stmap():
    return datetime.now().strftime("%H:%M:%S")


def PingFrontend(
    webhookImage: str,
    productSizes: list,
    title: str,
    price: str,
    url: str,
    pid: str,
    site: str,
):
    webhook = DiscordWebhook(url=os.getenv("WEBHOOK"), rate_limit_retry=True)
    embed = DiscordEmbed(
        title=title,
        url=url,
        color=2524623,
    )
    embed.set_thumbnail(url=webhookImage)
    embed.add_embed_field(name="Site", value=site, inline=True)

    embed.add_embed_field(name="Sizes", value=" ", inline=False)
    for size in productSizes:
        embed.add_embed_field(name="", value=size, inline=False)

    embed.add_embed_field(name="PID", value=pid, inline=True)

    embed.add_embed_field(name="Price", value="€{}".format(price), inline=True)
    embed.add_embed_field(
        name="MQT",
        value="[INSTOCK](https://quicktask.hellasaio.com/quicktask?product_id={}&siteId=2&size=random)".format(
            url
        ),
        inline=False,
    )

    embed.set_footer(
        text=f"EagleBot Front-End | {webhook_time_stmap()}",
        icon_url="https://media.discordapp.net/attachments/1039415918486376508/1065004102175699005/Screenshot_2022-11-17_at_09.47.01.png",
    )

    webhook.add_embed(embed)
    response = webhook.execute()
    if "<Response [405]>" in str(response):
        print("[ERROR] Webhook Incorrect")
    else:
        print("[INFO] Webhook Sent")


def PingServer(
    webhookImage: str,
    productSizes: list,
    title: str,
    price: str,
    url: str,
    pid: str,
    site: str,
):
    webhook = DiscordWebhook(url=os.getenv("WEBHOOK"), rate_limit_retry=True)
    embed = DiscordEmbed(
        title=title,
        url=url,
        color=2524623,
    )
    embed.set_thumbnail(url=webhookImage)
    embed.add_embed_field(name="Site", value=site, inline=True)
    embed.add_embed_field(name="Sizes", value=" ", inline=False)

    for size in productSizes:
        embed.add_embed_field(name="", value=size, inline=False)

    embed.add_embed_field(name="PID", value=pid, inline=True)

    embed.add_embed_field(name="Price", value="€{}".format(price), inline=True)
    embed.add_embed_field(
        name="MQT",
        value="[INSTOCK](https://quicktask.hellasaio.com/quicktask?product_id={}&siteId=2&size=random)".format(
            url
        ),
        inline=False,
    )

    embed.set_footer(
        text=f"EagleBot Back-End | {webhook_time_stmap()}",
        icon_url="https://media.discordapp.net/attachments/1039415918486376508/1065004102175699005/Screenshot_2022-11-17_at_09.47.01.png",
    )

    webhook.add_embed(embed)
    response = webhook.execute()
    if "<Response [405]>" in str(response):
        print("[ERROR] Webhook Incorrect")
    else:
        print("[INFO] Webhook Sent")
