## Simple Database Query 
Terdapat sebuah table "USER" yang memiliki 3 kolom: `ID`, `username`, `parent` 

Dimana:
- Kolom ID adalah Primary Key.
- Kolom username adalah Nama User.
- Kolom Parent adalah ID dari User yang menjadi Creator untuk User tertentu.

Example:
|----|----------|--------|
| id | username | parent |
|----|----------|--------|
| 1  | Ali      |    2   |
| 2  | Budi     |    0   |
| 3  | Cecep    |    1   |

Tuliskan SQL Query untuk mendapatkan data berisi:

|----|----------|-----------------|
| id | username | parent_username |
|----|----------|-----------------|
| 1  | Ali      |    Budi         |
| 2  | Budi     |    NULL         |
| 3  | Cecep    |    Ali          |

* Kolom Parent Username adalah Username berdasarkan value parent 
