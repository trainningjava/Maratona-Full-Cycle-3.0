from django.shortcuts import render
from maratona.models import Maratonas

# Create your views here.
def maratona_index(request):
    maratona = Maratonas.objects.all()
    context = {
        'maratona': maratona
    }
    return render(request, 'maratona_index.html', context)
