from typing import Optional
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class Book(BaseModel):
    title: str
    author: str
    price: int

@app.get("/")
def index():
    return {"hello": "world"}


@app.get("/books/{id}")
def get_book(id: int, q: Optional[str] = None):
    return {"id": id, "q": q}


@app.put("/books/{id}")
def update_item(id: int, book: Book):
    return {"title": book.title}