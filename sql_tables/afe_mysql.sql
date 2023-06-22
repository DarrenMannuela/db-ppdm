CREATE TABLE BIBLIOGRAPHY_TABLE_AFE (
    AFE_NUMBER INT,
    WORKSPACE_NAME VARCHAR(255),
    KKKS_NAME VARCHAR(255),
    WORKING_AREA VARCHAR(255),
    SUBMISSION_TYPE VARCHAR(255),
    DATA_TYPE VARCHAR(255) DEFAULT 'Bibliography',
    EMAIL VARCHAR(255),
    CONSTRAINT AFE_PK_BIBLIOGRAPHY_ERROR PRIMARY KEY (AFE_NUMBER),
    CONSTRAINT WORKSPACE_NAME_BIBLIOGRAPHY_UNIQUE UNIQUE (WORKSPACE_NAME),
    CONSTRAINT FK_BIBLIOGRAPHY_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
        REFERENCES SUBMISSION (SUBMISSION_TYPE),
    CONSTRAINT FK_BIBLIOGRAPHY_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
    );
    CREATE TABLE PRINT_WELL_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Print Well Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_PRINT_WELL_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_PRINT_WELL_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_PRINT_WELL_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_PRINT_WELL_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_MAPS_AND_TECHNICAL_DRAWING_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Maps And Technical Drawing',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_MAPS_AND_TECHNICAL_DRAWING_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_MAPS_AND_TECHNICAL_DRAWING_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_MAPS_AND_TECHNICAL_DRAWING_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_MAPS_AND_TECHNICAL_DRAWING_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Field Data Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_WELL_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Well Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_WELL_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_WELL_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_WELL_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_WELL_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_IMAGE_WELL_LOG_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Image Well Log',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_IMAGE_WELL_LOG_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_IMAGE_WELL_LOG_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_IMAGE_WELL_LOG_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_IMAGE_WELL_LOG_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_2D_3D_SEISMIC_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital 2D 3D Seismic Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_2D_3D_SEISMIC_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_2D_3D_SEISMIC_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_2D_3D_SEISMIC_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_2D_3D_SEISMIC_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Process Data Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WORKING_AREA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Working Area',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WORKING_AREA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WORKING_AREA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WORKING_AREA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WORKING_AREA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Non Seismic And Seismic Non Conventional Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_SECTION_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Section',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_SECTION_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_SECTION_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_SECTION_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_SECTION_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_DIGITAL_PROCESSED_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Digital Processed Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_DIGITAL_PROCESSED_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_DIGITAL_PROCESSED_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_DIGITAL_PROCESSED_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_DIGITAL_PROCESSED_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WELL_CORE_SAMPLE_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Well Core Sample',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WELL_CORE_SAMPLE_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WELL_CORE_SAMPLE_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WELL_CORE_SAMPLE_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WELL_CORE_SAMPLE_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE PRINT_TECHNICAL_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Print Technical Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_PRINT_TECHNICAL_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_PRINT_TECHNICAL_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_PRINT_TECHNICAL_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_PRINT_TECHNICAL_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Non Seismic And Seismic Non Conventional Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_NAVIGATION_DIGITAL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Navigation Digital Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_NAVIGATION_DIGITAL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_NAVIGATION_DIGITAL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_NAVIGATION_DIGITAL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_NAVIGATION_DIGITAL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Project File Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_PROJECT_FILE_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_PROJECT_FILE_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_PROJECT_FILE_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_PROJECT_FILE_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_FIELD_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Field Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_FIELD_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_FIELD_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_FIELD_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_FIELD_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WELL_LOGS_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Well Logs',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WELL_LOGS_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WELL_LOGS_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WELL_LOGS_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WELL_LOGS_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Non Seismic And Seismic Non Conventional Digital Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Navigation Data Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_3D_SEISMIC_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D 3D Seismic Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_3D_SEISMIC_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_3D_SEISMIC_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_3D_SEISMIC_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_3D_SEISMIC_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_WELL_LOG_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Well Log',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_WELL_LOG_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_WELL_LOG_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_WELL_LOG_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_WELL_LOG_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_PROCESSED_DIGITAL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Processed Digital Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_PROCESSED_DIGITAL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_PROCESSED_DIGITAL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_PROCESSED_DIGITAL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_PROCESSED_DIGITAL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_FIELD_DIGITAL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Field Digital Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_FIELD_DIGITAL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_FIELD_DIGITAL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_FIELD_DIGITAL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_FIELD_DIGITAL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Non Seismic And Seismic Non Conventional Data Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_SUMMARY_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Summary',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_SUMMARY_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_SUMMARY_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_SUMMARY_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_SUMMARY_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Navigation Data Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Non Seismic And Seismic Non Conventional Data Summary',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_SUMMARY_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Summary',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_SUMMARY_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_SUMMARY_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_SUMMARY_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_SUMMARY_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WELL_SAMPLES_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Well Samples',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WELL_SAMPLES_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WELL_SAMPLES_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WELL_SAMPLES_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WELL_SAMPLES_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE PRINT_MAPS_AND_TECHNICAL_DRAWING_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Print Maps And Technical Drawing',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_PRINT_MAPS_AND_TECHNICAL_DRAWING_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_PRINT_MAPS_AND_TECHNICAL_DRAWING_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_PRINT_MAPS_AND_TECHNICAL_DRAWING_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_PRINT_MAPS_AND_TECHNICAL_DRAWING_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_PROCESS_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Process Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_PROCESS_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_PROCESS_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_PROCESS_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_PROCESS_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '2D Seismic Navigation Digital Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WELL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Well Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WELL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WELL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WELL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WELL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Well Seismic Profile Data Stored In Media',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE FIELD_INFORMATION_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Field Information',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_FIELD_INFORMATION_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_FIELD_INFORMATION_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_FIELD_INFORMATION_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_FIELD_INFORMATION_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_TECHNICAL_REPORT_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Technical Report',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_TECHNICAL_REPORT_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_TECHNICAL_REPORT_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_TECHNICAL_REPORT_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_TECHNICAL_REPORT_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE BASIN_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Basin',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_BASIN_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_BASIN_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_BASIN_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_BASIN_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE WELL_SEISMIC_PROFILE_DIGITAL_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Well Seismic Profile Digital',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_WELL_SEISMIC_PROFILE_DIGITAL_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_WELL_SEISMIC_PROFILE_DIGITAL_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_WELL_SEISMIC_PROFILE_DIGITAL_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_WELL_SEISMIC_PROFILE_DIGITAL_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_PROJECT_FILE_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital Project File',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_PROJECT_FILE_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_PROJECT_FILE_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_PROJECT_FILE_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_PROJECT_FILE_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE OUTCROP_SAMPLE_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Outcrop Sample',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_OUTCROP_SAMPLE_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_OUTCROP_SAMPLE_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_OUTCROP_SAMPLE_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_OUTCROP_SAMPLE_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE T3D_SEISMIC_FIELD_DIGITAL_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT '3D Seismic Field Digital Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_T3D_SEISMIC_FIELD_DIGITAL_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_T3D_SEISMIC_FIELD_DIGITAL_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_T3D_SEISMIC_FIELD_DIGITAL_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_T3D_SEISMIC_FIELD_DIGITAL_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE SEISMIC_INTERPRETATION_DATA_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Seismic Interpretation Data',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_SEISMIC_INTERPRETATION_DATA_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_SEISMIC_INTERPRETATION_DATA_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_SEISMIC_INTERPRETATION_DATA_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_SEISMIC_INTERPRETATION_DATA_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                CREATE TABLE DIGITAL_2D_SEISMIC_SECTION_TABLE_AFE (
                AFE_NUMBER INT,
                WORKSPACE_NAME VARCHAR(255),
                KKKS_NAME VARCHAR(255),
                WORKING_AREA VARCHAR(255),
                SUBMISSION_TYPE VARCHAR(255),
                DATA_TYPE VARCHAR(255) DEFAULT 'Digital 2D Seismic Section',
                EMAIL VARCHAR(255),
                CONSTRAINT AFE_PK_DIGITAL_2D_SEISMIC_SECTION_ERROR PRIMARY KEY (AFE_NUMBER),
                CONSTRAINT WORKSPACE_NAME_DIGITAL_2D_SEISMIC_SECTION_UNIQUE UNIQUE (WORKSPACE_NAME),
                CONSTRAINT FK_DIGITAL_2D_SEISMIC_SECTION_SUBMISSION_TYPE FOREIGN KEY (SUBMISSION_TYPE)
                    REFERENCES SUBMISSION (SUBMISSION_TYPE),
                CONSTRAINT FK_DIGITAL_2D_SEISMIC_SECTION_USER_ID FOREIGN KEY (EMAIL) REFERENCES USERS(EMAIL)
                );
                