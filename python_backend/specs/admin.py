from django.contrib import admin

from .models import *

@admin.register(SparePartsUnitsFeatures)
class SparePartsUnitsAdmin(admin.ModelAdmin):
    list_display = ('id', 'feature_name')
    list_display_links = ('id', 'feature_name')
    list_per_page = 15
    search_fields = ('id', 'feature_name')
    search_help_text = 'Введите id или название'
    list_filter = ('id', 'feature_name')


@admin.register(FeatureValidators)
class SparePartsUnitsAdmin(admin.ModelAdmin):
    list_display = ('id', 'valid_feature_value')
    list_display_links = ('id', 'valid_feature_value')
    list_per_page = 15
    search_fields = ('id',)
    search_help_text = 'Введите id'
    list_filter = ('id',)


@admin.register(ProductFeatures)
class SparePartsUnitsAdmin(admin.ModelAdmin):
    list_display = ('id', 'value')
    list_display_links = ('id', 'value')
    list_per_page = 15
    search_fields = ('id',)
    search_help_text = 'Введите id'
    list_filter = ('id',)

