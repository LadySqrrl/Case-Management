DROP TABLE students;
DROP TABLE providers;
DROP TABLE services;

CREATE TABLE providers (
    providerID      INT AUTO_INCREMENT  NOT NULL    ,
    provider_name   VARCHAR(50)         NOT NULL    ,
    district        VARCHAR(100)                    ,
    building        VARCHAR(50)                     ,    
    CONSTRAINT pk_provider PRIMARY KEY (providerID)
);

CREATE TABLE students (
    studentID           INT AUTO_INCREMENT    NOT NULL      ,      
    student_last_name   VARCHAR(50)                         , 
    student_first_name  VARCHAR(50)                         ,
    dob                 DATE                                ,
    grade               INT                                 ,
    annual_due          DATE                                ,
    reeval_due          DATE                                ,
    weighting           DECIMAL(3, 2)                       ,
    case_manager        VARCHAR(50)                         ,
    roster_teacher      VARCHAR(50)                         ,
    CONSTRAINT pk_student PRIMARY KEY (studentID)           ,
    CONSTRAINT fk_roster_teacher FOREIGN KEY (roster_teacher) REFERENCES providers (providerID)
);

CREATE TABLE services (
    serviceID               INT AUTO_INCREMENT  NOT NULL    ,
    service_type            VARCHAR (5)                     ,
    providerID              VARCHAR(50)                     ,
    service_mins            INT                             ,
    service_frequency       VARCHAR(10)                     ,
    studentID               VARCHAR(100)                    ,
    CONSTRAINT pk_service PRIMARY KEY (serviceID),
    CONSTRAINT fk_provider FOREIGN KEY (providerID) REFERENCES providers (providerID),
    CONSTRAINT fk_student FOREIGN KEY (studentID) REFERENCES students (studentID)
);