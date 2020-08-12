from django.shortcuts import render

# Create your views here.
from django.shortcuts import render

def mensdesafio(request):
    return render(request, 'mensdesafio.html', {})