from django.db import models


class SiteLinks(models.Model):
    id = models.BigAutoField(primary_key=True)
    link = models.TextField(blank=True, null=True, verbose_name='Ссылка на сайт')

    class Meta:
        managed = False
        db_table = 'site_links'
        verbose_name = 'Ссылка на сайт'
        verbose_name_plural = 'Ссылки на сайт'
        ordering = ['id']

    def __str__(self):
        return f'{self.link}'


class Sliders(models.Model):
    id = models.BigAutoField(primary_key=True)
    main_text = models.TextField(blank=True, null=True, verbose_name='Главный текст')
    upper_text = models.TextField(blank=True, null=True, verbose_name='Текст выше')
    down_text = models.TextField(blank=True, null=True, verbose_name='Текст ниже')
    image = models.ImageField(upload_to="data/photos/sliders", null=True, verbose_name='Фотография')

    class Meta:
        managed = False
        db_table = 'sliders'
        verbose_name = 'Слайдер'
        verbose_name_plural = 'Слайдеры'
        ordering = ['id']

    def __str__(self):
        return f'{self.main_text}'


class MiniSliders(models.Model):
    id = models.BigAutoField(primary_key=True)
    upper_text = models.TextField(blank=True, null=True, verbose_name='Вверхний текст')
    marked_text = models.TextField(blank=True, null=True, verbose_name='Выделенный текст')
    main_text = models.TextField(blank=True, null=True, verbose_name='Главный текст')
    image = models.ImageField(upload_to="data/photos/labels", null=True, verbose_name='Фотография')

    class Meta:
        managed = False
        db_table = 'mini_sliders'
        verbose_name = 'Лейбл'
        verbose_name_plural = 'Лейбл'
        ordering = ['id']

    def __str__(self):
        return f'{self.upper_text}'


class News(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.TextField(blank=True, null=True, verbose_name='Заголовок')
    body = models.TextField(blank=True, null=True, verbose_name='Содержание')
    image = models.ImageField(upload_to="data/photos/news", null=True, verbose_name='Фотография')
    created_at = models.DateTimeField(blank=True, null=True, verbose_name='Дата создания')

    class Meta:
        managed = False
        db_table = 'news'
        verbose_name = 'Новость'
        verbose_name_plural = 'Новости'
        ordering = ['created_at']

    def __str__(self):
        return f'Заголовок - {self.title} | Дата публикации - {self.created_at}'


class Links(models.Model):
    id = models.BigAutoField(primary_key=True)
    link = models.TextField(blank=True, null=True, verbose_name='Ссылка на соц. сеть')

    class Meta:
        managed = False
        db_table = 'links'
        verbose_name = 'Ссылка на соц. сеть'
        verbose_name_plural = 'Ссылки на соц. сети'
        ordering = ['id']

    def __str__(self):
        return f'{self.link}'


class Addresses(models.Model):
    id = models.BigAutoField(primary_key=True)
    address = models.TextField(blank=True, null=True, verbose_name='Адрес')

    class Meta:
        managed = False
        db_table = 'addresses'
        verbose_name = 'Адресс'
        verbose_name_plural = 'Адреса'
        ordering = ['id']

    def __str__(self):
        return f'{self.address}'


class Services(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.TextField(blank=True, null=True, verbose_name='Название')
    body = models.TextField(blank=True, null=True, verbose_name='Содержание')
    image = models.ImageField(upload_to="data/photos/services", null=True, verbose_name='Фотография')

    class Meta:
        managed = False
        db_table = 'services'
        verbose_name = 'Услуга'
        verbose_name_plural = 'Услуги'

    def __str__(self):
        return f'Имя услуги - {self.title}'


class Abouts(models.Model):
    id = models.BigAutoField(primary_key=True)
    main_text = models.TextField(blank=True, null=True, verbose_name='Текст')
    image = models.ImageField(upload_to="data/photos/about", null=True, verbose_name='Фотография')

    class Meta:
        managed = False
        db_table = 'abouts'
        verbose_name = 'О компании'
        verbose_name_plural = 'О компании'

    def __str__(self):
        return f'{self.id}'


class Emails(models.Model):
    id = models.BigAutoField(primary_key=True)
    logo_image = models.ImageField(upload_to="data/photos/email_photo", verbose_name="Лого компании")
    cart_image = models.ImageField(upload_to="data/photos/email_photo", verbose_name="Картинка 'Корзины'")
    check_image = models.ImageField(upload_to="data/photos/email_photo", verbose_name="Картинка 'Галочки'")

    class Meta:
        managed = False
        db_table = 'emails'
        verbose_name = 'Содержимое Email'
        verbose_name_plural = 'Содержимое Email'

    def __str__(self):
        return f'{self.id},{self.logo_image}'


class EmailListToSends(models.Model):
    id = models.BigAutoField(primary_key=True)
    email = models.TextField(blank=True, null=True, verbose_name="Почта")
    can_to_send = models.BooleanField(blank=True, null=True, default=True, verbose_name="Статус")

    class Meta:
        managed = False
        db_table = 'email_list_to_sends'
        verbose_name = 'Email рассылку'
        verbose_name_plural = 'Email рассылкы'

    def __str__(self):
        return f'{self.email}'


class EmailListCategories(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.TextField(blank=True, null=True, verbose_name='Имя категории')
    emails_array = models.ManyToManyField(EmailListToSends, verbose_name='Почты')
    body = models.TextField(blank=True, null=True, verbose_name='Текст рассылки')

    class Meta:
        db_table = 'email_list_categories'
        verbose_name = 'Категория рассылки'
        verbose_name_plural = 'Категории рассылок'

    def __str__(self):
        return f'Имя категории - {self.title}'


class Metricts(models.Model):
    id = models.BigAutoField(primary_key=True)
    google_metric = models.TextField(blank=True, null=True, verbose_name='Google метрика')
    yandex_metric = models.TextField(blank=True, null=True, verbose_name='Yandex метрика')
    telegrambot_link = models.TextField(blank=True, null=True, verbose_name='Telegram бот')

    class Meta:
        managed = False
        db_table = 'metricts'
        verbose_name = 'Метрика'
        verbose_name_plural = 'Метрика'

    def __str__(self):
        return f'{self.id}'
