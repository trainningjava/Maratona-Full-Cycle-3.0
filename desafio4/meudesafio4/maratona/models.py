from django.db import models

# Create your models here.
class Maratonas(models.Model):
    aula = models.CharField(max_length=100)
