# -*- coding: utf-8 -*-
import os

from dotenv import load_dotenv
from langchain_core.messages import HumanMessage, SystemMessage
from langchain_gigachat.chat_models import GigaChat

load_dotenv("/")

giga = GigaChat(
    credentials=os.getenv("API_GIGACHAT"),
    verify_ssl_certs=False,
)

messages = [
    SystemMessage(
        content="Ты эмпатичный бот-психолог, который помогает пользователю решить его проблемы."
    )
]

while True:
    try:
        user_input = input("Пользователь: ")

        safe_input = ''.join(c for c in user_input if not (0xD800 <= ord(c) <= 0xDFFF))

        text = safe_input.strip()

        if text.lower() == "пока":
            break

        messages.append(HumanMessage(content=text))
        res = giga.invoke(messages)
        messages.append(res)
        print("GigaChat: ", res.content)

    except Exception as e:
        print(f"Произошла ошибка: {e}")
