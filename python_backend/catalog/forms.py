from django import forms

from .models import Categories, Products
class CategoriesForm(forms.ModelForm):

    class Meta:
        model = Categories
        fields = '__all__'

class ProductsForm(forms.ModelForm):

    class Meta:
        model = Products
        fields = '__all__'

