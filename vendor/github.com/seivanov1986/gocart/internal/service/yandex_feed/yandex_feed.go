package yandex_feed

import (
	"context"
	"encoding/xml"
	"strconv"
	"time"

	category2 "github.com/seivanov1986/gocart/internal/repository/category"
	"github.com/seivanov1986/gocart/internal/repository/product"
)

// https://yandex.ru/support/direct/feeds/requirements.html

type age struct {
	XMLName xml.Name `xml:"age"`
	Unit    string   `xml:"unit,attr"`
	Value   int      `xml:",chardata"`
}

type category struct {
	XMLName  xml.Name `xml:"category"`
	ID       int64    `xml:"id,attr"`
	ParentID int64    `xml:"parentId,attr"`
	Value    string   `xml:",chardata"`
}

type categories struct {
	Categories []category
}

type ymlCatalog struct {
	XMLName xml.Name `xml:"yml_catalog"`
	Date    string   `xml:"date,attr"`
	Shop    []shop   `xml:"shop"`
}

type offers struct {
	Offers []offer `xml:"offers"`
}

type shop struct {
	/* Короткое название магазина. В названии нельзя использовать слова, которые не относятся к
	наименованию магазина (например «лучший», «дешевый»), указывать номер телефона и т. п.
	Название магазина должно совпадать с фактическим названием, которое публикуется на сайте.
	Если требование не соблюдается, Яндекс.Маркет может самостоятельно изменить название
	без уведомления магазина. */
	Name string `xml:"name"`
	/* Полное наименование компании, владеющей магазином. Не публикуется. */
	Company string `xml:"company"`
	/* URL главной страницы магазина. Максимальная длина ссылки — 2048 символов.
	Допускаются кириллические ссылки. URL‑адрес формируется на основе стандарта RFC 3986. */
	Url string `xml:"url"`
	/* Система управления контентом, на основе которой работает магазин (CMS). */
	Platform *string `xml:"platform"`
	/* Версия CMS. */
	Version *string
	/* Наименование агентства, которое оказывает техническую поддержку магазину
	и отвечает за работоспособность сайта.*/
	Agency *string
	/* Контактный адрес разработчиков CMS или агентства, осуществляющего техподдержку. */
	Email *string
	/* Список курсов валют магазина. Должен быть перед списком предложений (offers). */
	Currencies *string
	/* Список категорий магазина. Обязательно должен быть перед списком предложений (offers). */
	Categories categories `xml:"categories"`
	/* Возможность курьерской доставки (по всем регионам, в которые доставляет магазин). Вы можете
	переопределить это значение в информации об отдельном товаре. Возможные значения:
	true — курьерская доставка есть, значение по умолчанию;
	false — курьерской доставки нет. */
	Delivery *string
	/* Стоимость и сроки курьерской доставки по региону, в котором находится магазин. */
	DeliveryOptions *string
	/* Возможность самовывоза из пунктов выдачи (во всех регионах, в которые доставляет магазин).
	Вы можете переопределить это значение в информации об отдельном товаре. Возможные значения:
	true — самовывоз есть, значение по умолчанию;
	false — самовывоза нет. */
	Pickup *string
	/* Стоимость и сроки самовывоза по региону, в котором находится магазин. */
	PickupOptions *string
	/* Автоматический расчет и показ скидок для всего прайс-листа.
	Обязательно должен быть перед списком предложений (offers).*/
	EnableAutoDiscounts *string
	/* Список предложений магазина. Каждое предложение описывается в отдельном элементе offer.
	Здесь не приводится список всех элементов, входящих в offer, так как он зависит от типа предложения.
	Для большинства категорий товаров подходят следующие типы описаний: */
	Offers offers `xml:"offers"`
	/* Подарки, которые не размещаются на Маркете (для акции «Подарок при покупке»). */
	Gifts *string
	/* Информация об акциях магазина. */
	Promos *string
	/* Возможность купить товар без предварительного заказа. Вы можете переопределить это значение
	в информации об отдельном товаре. Возможные значения:
	true — товары можно купить без предварительного заказа;
	false — товары нельзя купить без предварительного заказа. */
	Store *string
	/* Товары имеют отношение к удовлетворению сексуальных потребностей либо иным образом эксплуатируют
	интерес к сексу. Вы можете переопределить это значение в информации об отдельном товаре.
	Возможные значения: true, false. */
	Adult *string
}

type offer struct {
	XMLName   xml.Name `xml:"offer"`
	Id        int      `xml:"id,attr"`
	Available bool     `xml:"available,attr"`

	/* URL страницы товара. Обязательный элемент. */
	Url string `xml:"url"`
	/* Цена, по которой данный товар можно приобрести.
	Обязательный элемент для товарных дополнений.
	Рекомендуемый элемент для динамических объявлений и смарт-баннеров. */
	Price float64 `xml:"price"`
	/* Старая цена на товар, которая обязательно должна быть
	выше новой цены (price). */
	OldPrice *float32 `xml:"oldprice"`
	/* Код валюты (RUB, USD, UAH, KZT).
	Обязательный элемент, если есть элемент price. */
	CurrencyId string `xml:"currencyId"`
	/* Идентификатор категории товара, присвоенный рекламодателем
	(целое число не более 18 знаков). Товарное предложение может принадлежать
	только одной категории. Обязательный элемент. Элемент <offer> может содержать
	только один элемент <categoryId>. */
	CategoryId int `xml:"categoryId"`
	/* Ссылка на изображение. Обязательный элемент для смарт-баннеров
	и товарных дополнений. Добавьте несколько изображений
	в разных форматах (горизонтальное, вертикальное, квадратное),
	каждое в отдельном элементе <picture>. В смарт-баннере будет использовано
	то, которое лучше подходит по соотношению сторон. */
	Picture []string `xml:"picture"`
	/* Возможность купить товар в розничном магазине:
	true — товар можно купить в розничном магазине;
	false — возможность покупки в розничном магазине отсутствует. */
	Store bool `xml:"store"`
	/* Возможность самовывоза из пунктов выдачи:
	true — товар можно забрать самостоятельно;
	false — возможность самовывоза отсутствует. */
	Pickup bool `xml:"pickup"`
	/* Возможность курьерской доставки товара:
	true — возможна курьерская доставка;
	false — товар не может быть доставлен курьером. */
	Delivery bool `xml:"delivery"`
	/* Название товарного предложения. В названии упрощенного предложения
	рекомендуется указывать наименование и код производителя. Обязательный элемент.*/
	Name string `xml:"name"`
	/* Производитель */
	Vendor *string `xml:"vendor"`
	/* Код товара (указывается код производителя). */
	VendorCode *string `xml:"vendorCode"`
	/* Описание товарного предложения. */
	Description *string `xml:"description"`
	/* Информация о заказе:минимальной сумме заказа, минимальной партии товара,
	необходимости предоплаты; вариантах оплаты, описания акций и распродаж. */
	SalesNotes *string `xml:"sales_notes"`
	/* Наличие гарантии: true — товар имеет официальную гарантию;
	false — товар не имеет официальной гарантии. */
	ManufacturerWarranty bool `xml:"manufacturer_warranty"`
	/* Страна производства товара. */
	CountryOfOrigin *string `xml:"country_of_origin"`
	/* Возрастная категория товара. */
	Age *age `xml:"age"`
	/* Модель */
	Model *string `xml:"model"`
	/* Категория товара, в которой он должен быть размещен на Яндекс Маркете. */
	MarketCategory *string `xml:"market_category"`
	/* Товар относится к категории «для взрослых»: true — да; false — нет. */
	Adult bool `xml:"adult"`
	/* Возможность скачать товар:
	true — товар можно скачать;
	false — товар нельзя скачать. */
	Downloadable bool `xml:"downloadable"`
	/* Передача характеристик и параметров товара.
	В атрибуте name укажите название параметра (обязательно):
	Пол/Gender;
	Цвет/Colour или Color;
	Материал/Material;
	Размер/Size.
	Пример:
	<param name="Цвет">серый</param>
	<param name="Материал">алюминий</param> */
	Param []string `xml:"param"`
}

func (u *service) Generate(ctx context.Context, basePath, name, url, company string) string {
	offersList := []offer{}
	categoryList := []category{}

	categoryListFullInput := category2.CategoryListFullInput{
		Page: 0,
	}
	for {
		categories, err := u.hub.Category().ListFull(ctx, categoryListFullInput)
		if err != nil {
			break
		}

		if len(categories.List) == 0 {
			break
		}

		for _, v := range categories.List {
			category := category{
				ID:       v.ID,
				ParentID: v.ParentID,
				Value:    v.Name,
			}

			categoryList = append(categoryList, category)
		}

		categoryListFullInput.Page++
	}

	productListFullInput := product.ProductListFullInput{
		Page: 0,
	}
	for {
		productList, err := u.hub.Product().ListFull(ctx, productListFullInput)
		if err != nil {
			break
		}

		if len(productList.List) == 0 {
			break
		}

		for _, v := range productList.List {
			if v.Price == nil {
				continue
			}

			priceFloat, err := strconv.ParseFloat(*v.Price, 8)
			if err != nil {
				continue
			}

			offer := offer{
				Id:         int(v.Id),
				Name:       v.Name,
				Available:  true,
				Url:        basePath + v.Path + v.SfName,
				CurrencyId: "RUB",
				CategoryId: int(v.PtcIdCategory),
				Price:      priceFloat,
				Picture: []string{
					basePath + "/dynamic/400x340" + v.IPath + v.IName,
				},
				Store:                true,
				Pickup:               true,
				Delivery:             false,
				ManufacturerWarranty: true,
			}

			offersList = append(offersList, offer)
		}

		productListFullInput.Page++
	}

	ymlCatalog := ymlCatalog{}
	shop := shop{}
	shop.Name = name
	shop.Url = url
	shop.Company = company
	shop.Offers.Offers = offersList
	shop.Categories.Categories = categoryList
	ymlCatalog.Date = time.Now().Format("2006-01-02T15:04:05-07:00")
	ymlCatalog.Shop = append(ymlCatalog.Shop, shop)

	shopXml, _ := xml.MarshalIndent(ymlCatalog, "", " ")
	return "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" + string(shopXml) + "\n"
}
