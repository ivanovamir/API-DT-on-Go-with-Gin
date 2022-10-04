from django.contrib import admin

from .models import *
from .forms import *


class FeaturesInline(admin.TabularInline):
    model = Features
    extra = 1
    verbose_name = "Характеристики продукта"
    verbose_name_plural = "Характеристики продуктов"
    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False

    def has_change_permission(self, request, obj=None):
        return False


class OrderProductsInline(admin.TabularInline):
    model = OrderProducts
    extra = 1
    verbose_name = "Продукт заказа"
    verbose_name_plural = "Продукты заказа"
    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False

    def has_change_permission(self, request, obj=None):
        return False


class OrerInline(admin.TabularInline):
    model = Orders
    extra = 1
    verbose_name = "Заказ"
    verbose_name_plural = "Заказы"
    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False

    def has_change_permission(self, request, obj=None):
        return False


@admin.register(Categories)
class CategoriesAdmin(admin.ModelAdmin):
    list_display = ('id', 'title')
    list_display_links = ('id', 'title')
    list_per_page = 15
    search_fields = ('id', 'title')
    search_help_text = 'Введите id или название'
    form = CategoriesForm


@admin.register(SparePartsUnits)
class SparePartsUnitsAdmin(admin.ModelAdmin):
    list_display = ('id', 'title')
    list_display_links = ('id', 'title')
    list_per_page = 15
    search_fields = ('id', 'title')
    search_help_text = 'Введите id или название'


@admin.register(Products)
class ProductsAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', 'category', 'spare_parts_unit', 'price')
    list_display_links = ('id', 'title', 'price', 'category', 'spare_parts_unit')
    list_per_page = 15
    search_fields = ('id', 'title', 'price')
    search_help_text = 'Введите id,название или цену'
    list_filter = ('category', 'spare_parts_unit')
    form = ProductsForm
    change_form_template = 'custom_admin/change_form.html'
    inlines = (FeaturesInline,)


@admin.register(Orders)
class OrdersAdmin(admin.ModelAdmin):
    list_display = ('id', 'created_at',)
    list_display_links = ('id', 'created_at',)
    list_per_page = 15
    search_fields = ('id', 'created_at')
    search_help_text = 'Введите id или дату'
    inlines = (OrderProductsInline,)

    def has_add_permission(self, request, obj=None):
        return False

@admin.register(Feedbacks)
class FeedbacksAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'body', 'product', 'created_at')
    list_display_links = ('id', 'name', 'product')
    list_per_page = 15
    search_fields = ('id', 'name', 'created_at')
    search_help_text = 'Введите id, дату или имя автора'
    readonly_fields = ('product',)

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(FeedbackForms)
class FeedbackFormsAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'email', 'topic', 'message', 'created_at')
    list_display_links = ('id', 'name', 'email', 'topic', 'message', 'created_at')
    list_per_page = 15
    search_fields = ('id', 'created_at', 'name')
    search_help_text = 'Введите id, дату или имя автора'


@admin.register(Forms)
class FormsAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'email', 'phone', 'manager_phone', 'company_name',)
    list_display_links = ('id', 'name', 'email', 'phone', 'manager_phone', 'company_name',)
    list_per_page = 15
    search_fields = ('id', 'name', 'inn')
    search_help_text = 'Введите id заказа, имя или инн'
    inlines = (OrerInline,)


@admin.register(PriceLists)
class PriceListsAdmin(admin.ModelAdmin):
    list_display = ('id','price_1','price_2')
    list_display_links = ('id','price_1','price_2')

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(OrderProducts)
class OrderProductsAdmin(admin.ModelAdmin):
    list_display = ('order', 'product', 'count')
    list_display_links = ('order', 'product', 'count')
    list_per_page = 15
    search_fields = ('order__id',)
    search_help_text = 'Введите id заказа'
    readonly_fields = ('product',)


    def has_add_permission(self, request, obj=None):
        return False