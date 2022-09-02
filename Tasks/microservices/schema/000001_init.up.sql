CREATE TABLE USERS
(
    ID INT,
    PRIMARY KEY (id),
    name varchar(255) not null,
    age int
);

CREATE TABLE Friends
(
    ID1 INT,
    foreign key (ID1) references USERS(ID),
    ID2 INT,
    foreign key (ID2) references USERS(ID)
);