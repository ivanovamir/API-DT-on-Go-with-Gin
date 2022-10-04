from django import forms

from .models import FeatureValidators, SparePartsUnitsFeatures


class NewCategoryFeatureKeyForm(forms.ModelForm):

    class Meta:
        model = SparePartsUnitsFeatures
        fields = '__all__'


class FeatureValidatorForm(forms.ModelForm):

    class Meta:
        model = FeatureValidators
        fields = ['spare_parts_units']
