from django.db import models


class SparePartsUnitsFeatures(models.Model):
    id = models.BigAutoField(primary_key=True)
    feature_name = models.CharField(verbose_name='Имя ключа для категории', max_length=100)
    feature_filter_name = models.CharField(verbose_name='Имя для фильтра', max_length=100)
    unit = models.CharField(verbose_name='Единица измерения', max_length=100)
    spare_parts_units = models.ForeignKey("catalog.SparePartsUnits", models.DO_NOTHING, db_column='spare_parts_units', blank=True, null=True,verbose_name='Категория')

    class Meta:
        managed = False
        db_table = 'spare_parts_units_features'
        verbose_name = "Характеристики категории"
        verbose_name_plural = "Характеристики категории"

    def __str__(self):
        return f"{self.spare_parts_units.title} | {self.feature_name}"


class FeatureValidators(models.Model):
    id = models.BigAutoField(primary_key=True)
    valid_feature_value = models.CharField(verbose_name='Валидное значение', max_length=100)
    spare_parts_units = models.ForeignKey('catalog.SparePartsUnits', models.DO_NOTHING, db_column='spare_parts_units', blank=True, null=True, verbose_name='Подкатегория')
    feature_key = models.ForeignKey(SparePartsUnitsFeatures, models.DO_NOTHING, db_column='feature_key', blank=True, null=True, verbose_name='Ключ характеристики')

    class Meta:
        managed = False
        db_table = 'feature_validators'
        verbose_name = "Ключи характеристик"
        verbose_name_plural = "Ключи характеристик"

    def __str__(self):
        return f"{self.spare_parts_units.title} | " \
               f"{self.feature_key.feature_name} | " \
               f"{self.valid_feature_value} {self.feature_key.unit}"


class ProductFeatures(models.Model):
    id = models.BigAutoField(primary_key=True)
    value = models.CharField(verbose_name='Значение', max_length=100)
    product = models.ForeignKey('catalog.Products', models.DO_NOTHING, db_column='product', blank=True, null=True,verbose_name='Товар')
    feature = models.ForeignKey(SparePartsUnitsFeatures, models.DO_NOTHING, db_column='feature', blank=True, null=True, verbose_name='Характеристика')

    class Meta:
        db_table = 'product_features'
        verbose_name = "Характеристики продуктов"
        verbose_name_plural = "Характеристики продуктов"

    def __str__(self):
        return f"Товар - {self.product.title} | " \
               f"Характеристика - {self.feature.feature_name} | " \
               f"Значение - {self.text()}"

    def text(self):
        answer = self.value
        if self.feature.unit:
            answer += ' ' + self.feature.unit
        return answer