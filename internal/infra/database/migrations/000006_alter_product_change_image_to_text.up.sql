update product set image = null;

alter table product
alter column image type text;
