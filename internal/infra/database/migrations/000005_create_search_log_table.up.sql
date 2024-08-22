/* Entity SearchLog */
create table search_log (
	id SERIAL UNIQUE NOT NULL,
	user_id INT4 NOT NULL,
	field_key VARCHAR(255) NOT NULL,
	field_value VARCHAR(255) NOT NULL,
	search_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    foreign key (user_id) references users(id)
);
