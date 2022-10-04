from django.db import models


class OrderProducts(models.Model):
    id = models.BigAutoField(primary_key=True)
    order = models.ForeignKey('Orders', on_delete=models.CASCADE, blank=True, null=True, verbose_name='Номер заказа')
    product = models.ForeignKey('Products', models.DO_NOTHING, blank=True, null=True, verbose_name='Продукт')
    count = models.BigIntegerField(blank=True, null=True, verbose_name='Кол-во')

    class Meta:
        managed = False
        db_table = 'order_products'
        verbose_name = 'Продукт заказа'
        verbose_name_plural = 'Продукты заказа'
        ordering = ['-order']

    def __str__(self):
        return f'Номер заказа - {self.id} | Кол-во - {self.count}'


class Categories(models.Model):
    id = models.BigAutoField(primary_key=True, verbose_name = 'Id')
    title = models.TextField(blank=True, null=True, verbose_name = 'Имя категории')

    class Meta:
        managed = False
        db_table = 'categories'
        verbose_name = 'Категория'
        verbose_name_plural = 'Категории'
        ordering = ['id']

    def __str__(self):
        return f'{self.title}'


class SparePartsUnits(models.Model):
    id = models.BigAutoField(primary_key=True, verbose_name = 'Id')
    title = models.TextField(unique=True, verbose_name = 'Имя подкатегории')

    class Meta:
        managed = False
        db_table = 'spare_parts_units'
        verbose_name = 'Подкатегории'
        verbose_name_plural = 'Подкатегория'

    def __str__(self):
        return f'{self.title}'


class Feedbacks(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.TextField(blank=True, null=True, verbose_name = 'Имя')
    email = models.TextField(blank=True, null=True, verbose_name = 'Почта')
    body = models.TextField(blank=True, null=True, verbose_name = 'Отзыв')
    product = models.ForeignKey('Products', models.DO_NOTHING, blank=True, null=True, verbose_name = 'Продукт')
    created_at = models.DateTimeField(blank=True, null=True, verbose_name = 'Дата создания')

    class Meta:
        managed = False
        db_table = 'feedbacks'
        verbose_name = 'Отзыв'
        verbose_name_plural = 'Отзывы'
        ordering = ['created_at']

    def __str__(self):
        return f'Дата публикации - {self.created_at} | Товар - {self.product}'


class Params(models.Model):
    limit = models.BigIntegerField(blank=True, null=True)
    page = models.BigIntegerField(blank=True, null=True)
    cat_id = models.SmallIntegerField(blank=True, null=True)
    spare_parts_unit = models.SmallIntegerField(blank=True, null=True)
    
    class Meta:
        managed = False
        db_table = 'params'


class Products(models.Model):
    category = models.ForeignKey(Categories, models.DO_NOTHING, db_column='category', blank=True, null=True, verbose_name = 'Категория')
    spare_parts_unit = models.ForeignKey('SparePartsUnits', models.DO_NOTHING, db_column='spare_parts_unit', blank=True, null=True, verbose_name = 'Подкатегория')
    title = models.TextField(blank=True, null=True, verbose_name = 'Название')
    vendor_code =  models.TextField(blank=True, null=True, verbose_name = 'Артикул')
    description = models.TextField(blank=True, null=True, verbose_name = 'Описание')
    short_description = models.TextField(blank=True, null=True, verbose_name = 'Краткое описание')
    price = models.DecimalField(max_digits=15, decimal_places=2, blank=True, null=True, verbose_name='Цена')    
    image_original = models.ImageField(null=True, blank=True, upload_to="data/photos/original", verbose_name = 'Главная фотография')
    image_128 = models.ImageField(null=True, blank=True, upload_to="data/photos/128", verbose_name = 'Фотография 128px')
    image_432 = models.ImageField(null=True, blank=True, upload_to="data/photos/432", verbose_name = 'Фотография 432px')
    can_to_view = models.BooleanField(blank=True, null=True, verbose_name = 'Разрешение на публикацию')
    features_products = models.ManyToManyField("specs.ProductFeatures", blank=True, through='Features', related_name="features_products_reverse")

    class Meta:
        db_table = 'products'
        verbose_name = 'Товар'
        verbose_name_plural = 'Товары'
        ordering = ['id']

    def __str__(self):
        return f'Название- {self.title} | Категория - {self.category} | Подкатегория - {self.spare_parts_unit}'


class Features(models.Model):
    id = models.BigAutoField(primary_key=True)
    products = models.ForeignKey('Products', models.DO_NOTHING, blank=True, null=True)
    product_features = models.ForeignKey('specs.ProductFeatures', models.DO_NOTHING, blank=True, null=True)

    class Meta:
        db_table = 'features'


class FeedbackForms(models.Model):
    id = models.BigAutoField(primary_key=True)
    created_at = models.DateTimeField(blank=True, null=True, verbose_name = 'Дата создания')
    name = models.TextField(blank=True, null=True, verbose_name = 'Имя')
    email = models.TextField(blank=True, null=True, verbose_name = 'Email')
    topic = models.TextField(blank=True, null=True, verbose_name = 'Тема')
    message = models.TextField(blank=True, null=True, verbose_name = 'Содержание')

    class Meta:
        managed = False
        db_table = 'feedback_forms'
        verbose_name = 'Форма обратной связи'
        verbose_name_plural = 'Формы обратной связи'
        ordering = ['created_at']

    def __str__(self):
        return f'Тема - {self.topic} | Имя - {self.name}'


class Forms(models.Model):
    id = models.BigAutoField(primary_key=True, verbose_name='Номер заказа')
    company_name = models.TextField(blank=True, null=True, verbose_name='Компания')
    manager_phone = models.TextField(blank=True, null=True, verbose_name='Номер менеджера')
    manager_name = models.TextField(blank=True, null=True, verbose_name='Имя менеджера')
    inn = models.TextField(blank=True, null=True, verbose_name='ИНН')
    name = models.TextField(blank=True, null=True, verbose_name='ФИО')
    phone = models.TextField(blank=True, null=True, verbose_name='Телефон')
    email = models.TextField(blank=True, null=True, verbose_name='Почта')

    class Meta:
        managed = False
        db_table = 'forms'
        verbose_name = 'Форма покупателя'
        verbose_name_plural = 'Формы покупателей'
        ordering = ['id']

    def __str__(self):
        return f'Email - {self.email} | Имя - {self.name}'


class Orders(models.Model):
    id = models.BigAutoField(primary_key=True, verbose_name='Номер заказа')
    created_at = models.DateTimeField(blank=True, null=True, verbose_name='Дата создания')
    form = models.ForeignKey(Forms, models.DO_NOTHING, blank=True, null=True, verbose_name='Форма покупателя')
    note = models.TextField(blank=True, null=True, verbose_name='Примечание')

    class Meta:
        managed = False
        db_table = 'orders'
        verbose_name = 'Заказ'
        verbose_name_plural = 'Заказы'
        ordering = ['-created_at']

    def __str__(self):
        return f'Номер заказа - {self.id} | Дата создания - {self.created_at}'


class PriceLists(models.Model):
    id = models.BigAutoField(primary_key=True)
    price_1 = models.FileField(blank=True, upload_to="data/price_lists", verbose_name="Прайс лист 'Производство'")
    price_2 = models.FileField(blank=True, upload_to="data/price_lists", verbose_name="Прайс лист 'ТКР'")

    class Meta:
        managed = False
        db_table = 'price_lists'
        verbose_name = 'Прайс лист'
        verbose_name_plural = 'Прайс листы'

    def __str__(self):
        return f'{self.price_1}, {self.price_2}'