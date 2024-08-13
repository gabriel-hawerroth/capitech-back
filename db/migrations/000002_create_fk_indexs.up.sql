create index address_user_fk_idx on address(user_id);

create index product_category_fk_idx on product(category_id);

create index shopping_cart_user_fk_idx on shopping_cart(user_id);
create index shopping_cart_product_fk_idx on shopping_cart(product_id);

create index purchase_user_fk_idx on purchase(user_id);
create index purchase_address_fk_idx on purchase(address_id);

create index purchase_item_purchase_fk_idx on purchase_item(purchase_id);
create index purchase_item_product_fk_idx on purchase_item(product_id);
