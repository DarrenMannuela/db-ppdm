CREATE TABLE ROLES (
  ID NUMBER GENERATED ALWAYS AS IDENTITY,
  ROLE_NAME VARCHAR2(255) NOT NULL UNIQUE,
  PRIMARY KEY (ID)
);
