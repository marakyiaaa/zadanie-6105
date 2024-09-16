-- DROP TABLE IF EXISTS orders;
CREATE TABLE IF NOT EXISTS employee (
                          id SERIAL PRIMARY KEY,
                          username VARCHAR(50) UNIQUE NOT NULL,
                          first_name VARCHAR(50),
                          last_name VARCHAR(50),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE organization (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(100) NOT NULL,
                              description TEXT,
                              type organization_type,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
                                          id SERIAL PRIMARY KEY,
                                          organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
                                          user_id INT REFERENCES employee(id) ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS tenders (
                                       id SERIAL PRIMARY KEY,
                                       organization_id INT NOT NULL,
                                       status VARCHAR(255) NOT NULL,
                                       version INT NOT NULL,
                                       created_at TIMESTAMP NOT NULL,
                                       updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS proposals (
                                         id SERIAL PRIMARY KEY,
                                         tenderID INT NOT NULL,
                                         userID INT NOT NULL,
                                         description TEXT NOT NULL,
                                         price DECIMAL(10, 2) NOT NULL,
                                         created_at TIMESTAMP NOT NULL,
                                         updated_at TIMESTAMP NOT NULL,
                                         FOREIGN KEY (tenderID) REFERENCES tenders(id)
);