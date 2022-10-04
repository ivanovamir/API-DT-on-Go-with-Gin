from django.contrib import admin

from .models import *
from django import forms
from django_object_actions import DjangoObjectActions
from django.contrib import messages
from django.http import HttpResponseRedirect
from django.shortcuts import render
from django.urls import reverse, path
import requests


admin.site.site_url = None


@admin.register(SiteLinks)
class SiteLinksAdmin(admin.ModelAdmin):
    list_display = ('id', 'link')
    list_display_links = ('id', 'link')
    list_per_page = 15
    search_fields = ('id',)
    search_help_text = 'Введите id'

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(Sliders)
class SlidersAdmin(admin.ModelAdmin):
    list_display = ('id', 'main_text', 'upper_text', 'down_text')
    list_display_links = ('id', 'main_text', 'upper_text', 'down_text')
    list_per_page = 15
    search_fields = ('id',)
    search_help_text = 'Введите id'


@admin.register(News)
class NewsAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', 'created_at')
    list_display_links = ('id', 'title', 'created_at')
    list_per_page = 15
    search_fields = ('id', 'title', 'created_at')
    search_help_text = 'Введите id, заголовок или дату публикации'


@admin.register(Links)
class LinksAdmin(admin.ModelAdmin):
    list_display = ('id', 'link')
    list_display_links = ('id', 'link')
    list_per_page = 15
    search_fields = ('id', 'link')
    search_help_text = 'Введите id'

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(Addresses)
class AddressesAdmin(admin.ModelAdmin):
    list_display = ('id', 'address')
    list_display_links = ('id', 'address')
    list_per_page = 15
    search_fields = ('id',)
    search_help_text = 'Введите id,'

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False

@admin.register(Services)
class ServicesAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', )
    list_display_links = ('id', 'title', )

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(Abouts)
class AboutsAdmin(admin.ModelAdmin):
    list_display = ('id', 'main_text')
    list_display_links = ('id', 'main_text')

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.action(description='Сделать рассылку выбранных категорий')
def sender(modeladmin, request, queryset):
    emails_body = list(queryset.values_list('body', flat=True))
    emails_id_dict = list(queryset.values_list('emails_array', flat=True))
    post_data = {
    "key": "lydluxlhxlhx637447JfjfHzhg",
    "emails":list(emails_id_dict),
    "body":emails_body
    }
    response = requests.post('http://127.0.0.1:8080/email_sender', json=post_data)


@admin.register(Emails)
class EmailsAdmin(admin.ModelAdmin):
    list_display = ('id','logo_image','cart_image', 'check_image')
    list_display_links = ('id',)

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


class EmailListCsvsAdmin(forms.Form):
    csv_upload = forms.FileField()


@admin.register(EmailListToSends)
class EmailListToSendsAdmin(DjangoObjectActions, admin.ModelAdmin):
    list_display = ('id','email','can_to_send')
    list_display_links = ('id','email','can_to_send') 
    change_list_template = 'admin/cms/emaillisttosends/change_list.html'   


    def get_urls(self):
        urls = super().get_urls()
        new_urls = [path('upload-csv/', self.upload_csv),]
        return new_urls + urls
    

    def upload_csv(self, request):
        if request.method == "POST":
            csv_file = request.FILES["csv_upload"]
            
            if not csv_file.name.endswith('.csv'):
                messages.warning(request, 'Файл успешно загружен')
                return HttpResponseRedirect(request.path_info)
            
            file_data = csv_file.read().decode("ISO-8859-1")
            csv_data = file_data.split("\n")

            result = [line.rstrip() for line in csv_data]
            str_list_result = list(filter(None, result))
            for x in str_list_result:
                fields = x.split(",")
                created = EmailListToSends.objects.update_or_create(
                    email = fields[0],
                    ) 
            url = reverse('admin:index')
            return HttpResponseRedirect(url)

        form = EmailListCsvsAdmin()
        data = {"form": form}
        return render(request, "admin/cms/emaillisttosends/csv_upload.html", data)


@admin.register(EmailListCategories)
class EmailListCategoriesAdmin(admin.ModelAdmin):
    list_display = ('id', 'title')
    list_display_links = ('id', 'title')
    list_per_page = 15
    search_fields = ('id', 'title')
    search_help_text = 'Введите id или заголовок'
    filter_horizontal = ['emails_array']
    actions = [sender]


@admin.register(Metricts)
class MetrictsAdmin(admin.ModelAdmin):
    list_display = ('google_metric', 'yandex_metric', 'telegrambot_link')
    list_display_links = ('google_metric', 'yandex_metric', 'telegrambot_link')
    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False
