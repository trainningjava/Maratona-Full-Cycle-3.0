from django.urls import path
from . import views

urlpatterns = [
    path("", views.maratona_index, name="maratona_index"),
]