CREATE TABLE USERS (
  ID INT AUTO_INCREMENT,
  FIRST_NAME VARCHAR(255) NOT NULL,
  LAST_NAME VARCHAR(255) NOT NULL,
  EMAIL VARCHAR(255) NOT NULL,
  PASSWORD VARCHAR(255) NOT NULL,
  DATE_JOINED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  EXPIRED_DATE DATE NOT NULL,
  ROLE_NAME VARCHAR(255) NOT NULL,
  PROFILE_PICTURE LONGTEXT,
  PRIMARY KEY (ID),
  UNIQUE (EMAIL),
  CONSTRAINT FK_CUSTOMER_ROLE FOREIGN KEY (ROLE_NAME)
    REFERENCES ROLES (ROLE_NAME)
);
