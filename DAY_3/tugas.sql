-- Perintah Create
create table customers(
	id int not null,
	customer_name char(50) not null,
	primary key(id)
);

create table products(
	id int not null,
	name char(50) not null,
	primary key(id)
);

create table Orders(
	order_id int not null,
	customer_id int not null,
	product_id int not null,
	order_date date not null,
	total float not null,
	
	primary key (order_id),
	foreign key (customer_id) references customers (id),
	foreign key (product_id) references products (id)
);

table customers;
table products;
table orders ;

select * from customers,products,Orders;

-- Perintah Insert
insert into public.customers (id,customer_name)
values 
	(101,'adit'),
	(102,'diva'),
	(103,'djikstra'),
	(104,'firman');

insert into public.products (id,name)
values 
	(201,'Laptop'),
	(202,'PS'),
	(203,'TV'),
	(204,'HP');
	

insert into public.Orders (order_id,customer_id,product_id,order_date,total)
values 
	(1,101,201,'2022-10-26',1),
	(2,102,202,'2022-10-27',2),
	(3,103,203,'2022-10-28',3),
	(4,104,204,'2022-10-29',4),
	(5,101,202,'2022-10-30',1),
	(6,102,203,'2022-10-31',2),
	(7,103,204,'2022-11-1',5),
	(8,104,201,'2022-11-2',3),
	(9,102,204,'2022-11-3',2),
	(10,101,204,'2022-11-4',1);

-- Perintah Update
update public.orders 
set total = 10
where order_id =7;

update public.customers  
set customer_name = 'Alex'
where id = 104;

update public.products  
set name = 'PC'
where id =204;

-- Perintah delete 
delete 
from public.orders 
where order_id =10;

-- Perintah dibawah error karena terdapat row pada table orders yang menggunakan id tersebut sebagai foreign key
delete 
from public.customers  
where id=102;

-- Perintah dibawah error karena terdapat row pada table orders yang menggunakan id tersebut sebagai foreign key
delete 
from public.products  
where id=201;


-- Perintah Join
select * 
from orders join customers on customer_id = id; 


select * 
from orders join products on product_id = id; 




