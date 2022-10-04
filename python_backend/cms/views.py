from django.shortcuts import redirect

from .models import Metricts


def CmsGoogleMetricLink(request):
    link = Metricts.objects.get(id=1)
    field_name_val = getattr(link, "google_metric")
    return redirect (field_name_val)

def CmsYandexMetricLink(request):
    link = Metricts.objects.get(id=1)
    field_name_val = getattr(link, "yandex_metric")
    return redirect (field_name_val)

def CmsTelegramBotLink(request):
    link = Metricts.objects.get(id=1)
    field_name_val = getattr(link, "telegrambot_link")
    return redirect (field_name_val)

