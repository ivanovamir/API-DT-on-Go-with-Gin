from collections import defaultdict

from django.contrib import messages
from django.shortcuts import render
from django.views.generic import View
from django.http import HttpResponseRedirect, JsonResponse

from .models import SparePartsUnitsFeatures, FeatureValidators, ProductFeatures
from .forms import NewCategoryFeatureKeyForm
from catalog.models import Products, SparePartsUnits


class BaseSpecView(View):
    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            return render(request, 'product_features.html', {})
        else:
            return HttpResponseRedirect('/')


class CreateNewFeature(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            form = NewCategoryFeatureKeyForm(request.POST or None)
            context = {'form': form}
            return render(request, 'new_feature.html', context)
        else:
            return HttpResponseRedirect('/')

    def post(self, request, *args, **kwargs):
        if request.user.is_superuser:
            form = NewCategoryFeatureKeyForm(request.POST or None)
            if form.is_valid():
                new_category_feature_key = form.save(commit=False)
                new_category_feature_key.spare_parts_units = form.cleaned_data['spare_parts_units']
                new_category_feature_key.feature_name = form.cleaned_data['feature_name']
                new_category_feature_key.save()

                messages.add_message(
                    request, messages.SUCCESS,
                   f'Новая характеристика для категории {new_category_feature_key.spare_parts_units.title} создана'
                )
            return HttpResponseRedirect('/product-specs/')
        else:
            return HttpResponseRedirect('/')


class CreateNewFeatureValidator(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            categories = SparePartsUnits.objects.all()
            context = {'categories': categories}
            return render(request, 'new_validator.html', context)
        else:
            return HttpResponseRedirect('/')


class FeatureChoiceView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            option = '<option value="{value}">{option_name}</option>'
            html_select = """
                <select class="form-select" name="feature-validators" id="feature-validators-id" aria-label="Default select example">
                    <option selected>---</option>
                    {result}
                </select>
                        """
            feature_key_qs = SparePartsUnitsFeatures.objects.filter(
                spare_parts_units_id=int(request.GET.get('category_id'))
            )
            res_string = ""
            for item in feature_key_qs:
                res_string += option.format(value=item.feature_name, option_name=item.feature_name)
            html_select = html_select.format(result=res_string)
            return JsonResponse({"result": html_select, "value": int(request.GET.get('category_id'))})
        else:
            return HttpResponseRedirect('/')


class CreateFeatureView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            spare_parts_units_id = request.GET.get('category_id')
            feature_name = request.GET.get('feature_name')
            value = request.GET.get('feature_value').strip(" ")
            print(value)
            spare_parts_units = SparePartsUnits.objects.get(id=int(spare_parts_units_id))
            feature = SparePartsUnitsFeatures.objects.get(spare_parts_units=spare_parts_units, feature_name=feature_name)
            existed_object, created = FeatureValidators.objects.get_or_create(
                spare_parts_units=spare_parts_units,
                feature_key=feature,
                valid_feature_value=value
            )
            if not created:
                return JsonResponse({
                    "error": f"Значение '{value}' уже существует."
                })
            messages.add_message(
                request, messages.SUCCESS,
                f'Значение "{value}" для характеристики '
                f'"{feature.feature_name}" в категории {spare_parts_units.title} успешно создано'
            )
            return JsonResponse({'result': 'ok'})
        else:
            return HttpResponseRedirect('/')


class NewProductFeatureView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            categories = SparePartsUnits.objects.all()
            context = {'categories': categories}
            return render(request, 'new_product_feature.html', context)
        else:
            return HttpResponseRedirect('/')


class SearchProductAjaxView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            query = request.GET.get('query')
            category_id = request.GET.get('category_id')
            category = SparePartsUnits.objects.get(id=int(category_id))
            products = list(Products.objects.filter(
                spare_parts_unit=category,
                title__icontains=query
            ).values())
            return JsonResponse({"result": products})
        else:
            return HttpResponseRedirect('/')

# ERROR_HERE
class AttachNewFeatureToProduct(View): 

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            res_string = ""
            product = Products.objects.get(id=int(request.GET.get('product_id')))
            existing_features = list(set([item.feature.feature_name for item in product.features_products.all()]))
            print(existing_features)
            category_features = SparePartsUnitsFeatures.objects.filter(
                spare_parts_units=product.spare_parts_unit
            ).exclude(feature_name__in=existing_features)
            option = '<option value="{value}">{option_name}</option>'
            html_select = """
                <select class="form-select" name="product-category-features" id="product-category-features-id" aria-label="Default select example">
                    <option selected>---</option>
                    {result}
                </select>
                        """
            for item in category_features:
                res_string += option.format(value=item.spare_parts_units.id, option_name=item.feature_name)
            html_select = html_select.format(result=res_string)
            return JsonResponse({"features": html_select})
        else:
            return HttpResponseRedirect('/')


class ProductFeatureChoicesAjaxView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            res_string = ""
            category = SparePartsUnits.objects.get(id=int(request.GET.get('category_id')))
            feature_key = SparePartsUnitsFeatures.objects.get(
                spare_parts_units=category,
                feature_name=request.GET.get('product_feature_name')
            )
            validators_qs = FeatureValidators.objects.filter(
                spare_parts_units=category,
                feature_key=feature_key
            )
            option = '<option value="{value}">{option_name}</option>'
            html_select = """
                <select class="form-select" name="product-category-features-choices" id="product-category-features-choices-id" aria-label="Default select example">
                    <option selected>---</option>
                    {result}
                </select>
                        """
            for item in validators_qs:
                res_string += option.format(value=item.id, option_name=item.valid_feature_value)
            html_select = html_select.format(result=res_string)
            return JsonResponse({"features": html_select})
        else:
            return HttpResponseRedirect('/')


class CreateNewProductFeatureAjaxView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            product = Products.objects.get(title=request.GET.get('product'))
            category_feature = SparePartsUnitsFeatures.objects.get(
                spare_parts_units=product.spare_parts_unit,
                feature_name=request.GET.get('category_feature')
            )
            value = request.GET.get('value')
            feature = ProductFeatures.objects.create(
                feature=category_feature,
                product=product,
                value=value
            )
            product.features_products.add(feature)
            messages.add_message(
                request, messages.SUCCESS,
                f'Характеристика для товара {product.title} успешно создана'
            )
            return JsonResponse({"OK": "OK"})
        else:
            return HttpResponseRedirect('/')


class UpdateProductFeaturesView(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            categories = SparePartsUnits.objects.all()
            context = {'categories': categories}
            return render(request, 'update_product_features.html', context)
        else:
            return HttpResponseRedirect('/')

# ERROR_HERE    
class ShowProductFeaturesForUpdate(View):

    def get(self, request, *args, **kwargs):
        if request.user.is_superuser:
            product = Products.objects.get(id=int(request.GET.get('product_id')))
            features_values_qs = product.features_products.all()
            head = """
            <hr>
                <div class="row">
                    <div class="col-md-4">
                        <h4 class="text-center">Характеристика</h4>
                    </div>
                    <div class="col-md-4">
                        <h4 class="text-center">Текущее значение</h4>
                    </div>
                    <div class="col-md-4">
                        <h4 class="text-center">Новое значение</h4>
                    </div>
                </div>
            <div class='row'>{}</div>
            <div class="row">
            <hr>
            <div class="col-md-4">
            </div>
            <div class="col-md-4">
                <p class='text-center'><button class="btn btn-success" id="save-updated-features">Сохранить</button></p> 
            </div>
            <div class="col-md-4">
            </div>
            </div>
            """
            option = '<option value="{value}">{option_name}</option>'
            select_values = """
                <select class="form-select" name="feature-value" id="feature-value" aria-label="Default select example">
                    <option selected>---</option>
                    {result}
                </select>
                        """
            mid_res = ""
            select_different_values_dict = defaultdict(list)
            for item in features_values_qs:
                fv_qs = FeatureValidators.objects.filter(
                    spare_parts_units=item.product.spare_parts_unit,
                    feature_key=item.feature
                ).values()
                for fv in fv_qs:
                    if fv['valid_feature_value'] == item.value:
                        pass
                    else:
                        select_different_values_dict[fv['feature_key_id']].append(fv['valid_feature_value'])
                feature_field = '<input type="text" class="form-control" id="{id}" value="{value}" disabled/>'
                current_feature_value = """
                <div class='col-md-4 feature-current-value' style='margin-top:10px; margin-bottom:10px;'>{}</div>
                                        """
                body_feature_field = """
                <div class='col-md-4 feature-name' style='margin-top:10px; margin-bottom:10px;'>{}</div>
                                    """
                body_feature_field_value = """
                <div class='col-md-4 feature-new-value' style='margin-top:10px; margin-bottom:10px;'>{}</div>
                """
                body_feature_field = body_feature_field.format(feature_field.format(id=item.feature.id, value=item.feature.feature_name))
                current_feature_value_mid_res = ""
                for item_ in select_different_values_dict[item.feature.id]:
                    current_feature_value_mid_res += option.format(value=item.feature.id, option_name=item_)
                body_feature_field_value = body_feature_field_value.format(
                    select_values.format(item.feature.id, result=current_feature_value_mid_res)
                )
                current_feature_value = current_feature_value.format(feature_field.format(id=item.feature.id, value=item.value))
                m = body_feature_field + current_feature_value + body_feature_field_value
                mid_res += m
            result = head.format(mid_res)
            return JsonResponse({"result": result})
        else:
            return HttpResponseRedirect('/')


class UpdateProductFeaturesAjaxView(View):

    def post(self, request, *args, **kwargs):
        if request.user.is_superuser:
            features_names = request.POST.getlist('features_names')
            features_current_values = request.POST.getlist('features_current_values')
            new_feature_values = request.POST.getlist('new_feature_values')
            data_for_update = [{'feature_name': name, 'current_value': curr_val, 'new_value': new_val} for name, curr_val, new_val
                               in zip(features_names, features_current_values, new_feature_values)]
            product = Products.objects.get(title=request.POST.get('product'))
            for item in product.features_products.all():
                for item_for_update in data_for_update:
                    if item.feature.feature_name == item_for_update['feature_name']:
                        if item.value != item_for_update['new_value'] and item_for_update['new_value'] != '---':
                            cf = SparePartsUnitsFeatures.objects.get(
                                spare_parts_units=product.spare_parts_unit,
                                feature_name=item_for_update['feature_name']
                            )
                            item.value = FeatureValidators.objects.get(
                                spare_parts_units=product.spare_parts_unit,
                                feature_key=cf,
                                valid_feature_value=item_for_update['new_value']
                            ).valid_feature_value
                            item.save()
            messages.add_message(
                request, messages.SUCCESS,
                f'Значения характеристик для товара {product.title} успешно обновлены'
            )
            return JsonResponse({"result": "ok"})
        else:
            return HttpResponseRedirect('/')