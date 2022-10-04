from django.urls import path

from cms.views import CmsTelegramBotLink, CmsGoogleMetricLink, CmsYandexMetricLink


urlpatterns = [
    path('telegram_bot/', CmsTelegramBotLink, name='telegram_bot'),
    path('google_metric/', CmsGoogleMetricLink, name='google_metric'),
    path('yandex_metric/', CmsYandexMetricLink, name='yandex_metric'),

]
