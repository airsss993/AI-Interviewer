from fastapi import FastAPI
from pydantic import BaseModel
import os
from dotenv import load_dotenv
from langchain_core.messages import HumanMessage, SystemMessage
from langchain_gigachat.chat_models import GigaChat


class Prompt(BaseModel):
    text: str


giga = GigaChat(
    credentials=os.getenv("API_GIGACHAT"),
    verify_ssl_certs=False,
)

app = FastAPI()

load_dotenv(dotenv_path="../../.env")

system_prompt = """
    Ты — искусственный интеллект, созданный для проведения технического интервью с пользователями.
    Твоя цель — задавать технические вопросы, анализировать ответы и оценивать их с точки зрения логики, точности, глубины и качества решения задач.

    Ты должен:
    1. Задавать технические вопросы по программированию, алгоритмам, структурам данных и другим соответствующим темам.
    2. Отвечать только на сообщения, которые являются ответами на твои вопросы или связаны с интервью.
    3. Если пользователь прислал сообщение, которое не является ответом на твой вопрос, не отвечать на него и напомнить пользователю о необходимости ответить на вопрос.
    4. Оценивать ответы по шкале от 1 до 10 в зависимости от их полноты, точности, правильности решения задачи и логики.
    5. В случае неполных или неясных ответов задать уточняющий вопрос для получения более точной информации.
    6. Вопросы и ответы должны быть связаны с техническими аспектами.
    7. Вести беседу в дружелюбном, но профессиональном тоне.
    8. По завершении интервью дать рекомендации по улучшению технических навыков, основываясь на оценках.
    """

messages = [
    SystemMessage(content=system_prompt)
]


@app.post("/chat")
async def create_item(prompt: Prompt):
    user_input = prompt.text
    safe_input = ''.join(c for c in user_input if not (0xD800 <= ord(c) <= 0xDFFF))
    text = safe_input.strip()
    messages.append(HumanMessage(content=text))
    res = giga.invoke(messages)
    messages.append(res)
    return {"response": res.content}
