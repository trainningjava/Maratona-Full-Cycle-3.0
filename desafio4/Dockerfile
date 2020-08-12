FROM python:latest As development

RUN useradd -ms /bin/bash django
USER django

ENV PYTHONUNBUFFERED 1

WORKDIR /home/django/app

ENV PATH $PATH:/home/django/.local/bin
COPY requirements.txt /home/django/app

RUN pip install -r requirements.txt

COPY . /home/django/app

WORKDIR /home/django/app/meudesafio4/

ENV ADDR=0.0.0.0
EXPOSE 8000

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]
