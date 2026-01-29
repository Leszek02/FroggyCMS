/*  ---------------------------------------------------
    Theme Name: Florist
    Description: Florist - E-commerce Template
    Author: Colorib
    Author URI: https://www.colorib.com/
    Version: 1.0
    Created: Colorib
---------------------------------------------------------  */

'use strict';

(function ($) {

    /*------------------
        Preloader
    --------------------*/
    $(window).on('load', function () {
        $(".loader").fadeOut();
        $("#preloder").delay(200).fadeOut("slow");

        /*------------------
            Gallery filter
        --------------------*/
        $('.filter__controls li').on('click', function () {
            $('.filter__controls li').removeClass('active');
            $(this).addClass('active');
        });
        if ($('.product__filter').length > 0) {
            var containerEl = document.querySelector('.product__filter');
            var mixer = mixitup(containerEl);
        }
    });

    /*------------------
        Background Set
    --------------------*/
    $('.set-bg').each(function () {
        var bg = $(this).data('setbg');
        $(this).css('background-image', 'url(' + bg + ')');
    });

    //Canvas Menu
    $(".canvas__open").on('click', function () {
        $(".offcanvas-menu-wrapper").addClass("active");
        $(".offcanvas-menu-overlay").addClass("active");
    });

    $(".offcanvas-menu-overlay").on('click', function () {
        $(".offcanvas-menu-wrapper").removeClass("active");
        $(".offcanvas-menu-overlay").removeClass("active");
    });

    //Search Switch
    $('.search-switch').on('click', function () {
        $('.search-model').fadeIn(400);
    });

    $('.search-close-switch').on('click', function () {
        $('.search-model').fadeOut(400, function () {
            $('#search-input').val('');
        });
    });

    /*------------------
		Navigation
	--------------------*/
    $(".mobile-menu").slicknav({
        prependTo: '#mobile-menu-wrap',
        allowParentLinks: true
    });

    /*-----------------------
        Hero Slider
    ------------------------*/
    $(".hero__slider").owlCarousel({
        loop: true,
        margin: 0,
        items: 1,
        dots: true,
        nav: false,
        smartSpeed: 1200,
        autoHeight: false,
        animateOut: 'fadeOut',
        animateIn: 'fadeIn',
        autoplay: true,
        mouseDrag: true
    });

    /*-----------------------
        Testimonial Slider
    ------------------------*/
    $(".testimonial__slider").owlCarousel({
        loop: true,
        margin: 0,
        items: 1,
        dots: true,
        nav: true,
        navText: ["<i class='fa fa-angle-left'><i/>", "<i class='fa fa-angle-right'><i/>"],
        smartSpeed: 1200,
        autoHeight: false,
        animateOut: 'fadeOut',
        animateIn: 'fadeIn',
        autoplay: true
    });

    /*--------------------------
        Select
    ----------------------------*/
    $('select:not([name="delivery_method"]):not([name="payment_method"])').niceSelect();

    /*------------------
		Magnific
	--------------------*/
    $('.video-popup').magnificPopup({
        type: 'iframe'
    });

    /*------------------
		Single Product
	--------------------*/
    $('.pt__item img').on('click', function () {
        var imgurl = $(this).data('imgbigurl');
        var bigImg = $('.product__details__big__pic').attr('src');
        if (imgurl != bigImg) {
            $('.product__details__big__pic').attr({
                src: imgurl
            });
        }
    });

    $(".nice-scroll").niceScroll({
        cursorborder:"",
        cursorcolor:"#dddddd",
        boxzoom:false,
        cursorwidth: 5,
        background: 'rgba(0, 0, 0, 0.2)',
        cursorborderradius:50,
        horizrailenabled: false
    });

    /*-------------------
		Quantity change
	--------------------- */
    var proQty = $('.cart__quantity .pro-qty')
    proQty.on('click', '.qtybtn', function () {
        var $button = $(this);
        var $input = $button.parent().find('input');
        var oldValue = parseFloat($input.val());
        
        if ($button.hasClass('inc')) {
            var newVal = oldValue + 1;
        } else {
            if (oldValue > 1) {
                var newVal = oldValue - 1;
            } else {
                newVal = 1;
            }
        }
        
        $input.val(newVal);
        updateItemTotal($input);
        updateCartTotals();
        
        saveQuantityChange($input);
    });

    function saveQuantityChange($input) {
        var productId = $input.closest('tr').data('product-id');
        var quantity = parseInt($input.val());
        
        $.ajax({
            url: '/carts',
            method: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify({ newQuantity: quantity, product_id: productId }),
            success: function(response) {
                console.log('Cart updated');
            },
            error: function(xhr) {
                alert('Failed to update cart');
            }
        });
    }

    function updateItemTotal($input) {
        var quantity = parseInt($input.val());

        var $row = $input.closest('tr');

        var priceText = $row.find('.cart__item__text span')
            .text()
            .replace('$', '');

        var price = parseFloat(priceText);
        var itemTotal = (price * quantity).toFixed(2);

        $row.find('.cart__price').text('$' + itemTotal);
    }

    function updateCartTotals() {
        var cartTotal = 0;

        $('.cart__price').each(function () {
            cartTotal += parseFloat($(this).text().replace('$', ''));
        });
        console.log(cartTotal)
        $('.cart__total li span').text('$' + cartTotal.toFixed(2));
}


    var prodQty = $('.product__details__btns .pro-qty')
    prodQty.on('click', '.qtybtn', function () {
            var $button = $(this);
            var $input = $button.parent().find('input');
            var oldValue = parseFloat($input.val());
            
            if ($button.hasClass('inc')) {
                var newVal = oldValue + 1;
            } else {
                if (oldValue > 1) {
                    var newVal = oldValue - 1;
                } else {
                    newVal = 1;
                }
            }
            
            $input.val(newVal);
        });
    /*------------------
        Achieve Counter
    --------------------*/
    $('.c_num').each(function () {
        $(this).prop('Counter', 0).animate({
            Counter: $(this).text()
        }, {
            duration: 4000,
            easing: 'swing',
            step: function (now) {
                $(this).text(Math.ceil(now));
            }
        });
    });

})(jQuery);