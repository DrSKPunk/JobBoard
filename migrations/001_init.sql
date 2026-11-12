CREATE TABLE Polzovateli (
    ID_polzovatelya SERIAL PRIMARY KEY,
    Imya VARCHAR(100),
    Familiya VARCHAR(100),
    Email VARCHAR(255) UNIQUE NOT NULL,
    Parol VARCHAR(255) NOT NULL
);

CREATE TABLE Kompanii (
    ID_kompanii SERIAL PRIMARY KEY,
    Nazvanie VARCHAR(255) NOT NULL,
    Opisanie TEXT,
    Gorod VARCHAR(100),
    Telefon VARCHAR(50),
    Email VARCHAR(255)
);

CREATE TABLE Vakansii (
    ID_vakansii SERIAL PRIMARY KEY,
    ID_kompanii INTEGER REFERENCES Kompanii(ID_kompanii) ON DELETE CASCADE,
    Nazvanie VARCHAR(255) NOT NULL,
    Opisanie TEXT,
    Zarplata INTEGER,
    Gorod VARCHAR(100),
    Data_publikacii DATE
);

CREATE TABLE Rezume (
    ID_rezume SERIAL PRIMARY KEY,
    ID_polzovatelya INTEGER REFERENCES Polzovateli(ID_polzovatelya) ON DELETE CASCADE,
    Zhelaemaya_dolzhnost VARCHAR(255),
    Obrazovanie TEXT,
    Opyt_raboty TEXT,
    O_sebe TEXT,
    Gorod VARCHAR(100)
);

CREATE TABLE Navyki (
    ID_navyka SERIAL PRIMARY KEY,
    Nazvanie VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Rezume_Navyki (
    ID_rezume INTEGER REFERENCES Rezume(ID_rezume) ON DELETE CASCADE,
    ID_navyka INTEGER REFERENCES Navyki(ID_navyka) ON DELETE CASCADE,
    PRIMARY KEY (ID_rezume, ID_navyka)
);

CREATE TABLE Otkliki (
    ID_otklika SERIAL PRIMARY KEY,
    ID_rezume INTEGER REFERENCES Rezume(ID_rezume) ON DELETE CASCADE,
    ID_vakansii INTEGER REFERENCES Vakansii(ID_vakansii) ON DELETE CASCADE,
    Data_otklika TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Status VARCHAR(50) DEFAULT 'new'
);