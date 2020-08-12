from django.urls import path

from mensdesafio import views

urlpatterns = [path("", views.mensdesafio, name="mensdesafio")]