## quiz-3-projects

base url: https://quiz-3-projects-production.up.railway.app

### routing :
test pada POSTMAN

bangun-datar: 
- persegi (GET): /persegi?sisi=(number)&hitung=(luas|keliling)
  - Response :
    ```
    {
      "message": "string",
      "data": {
        "keliling": 0,
        "luas": 0,
        "sisi": 0
      }
    }
    ```
- persegi panjang (GET): /persegi-panjang?panjang=(number)&lebar=(number)&hitung=(luas|keliling)
  - Response :
    ```
    {
      "message": "string",
      "data": {
        "keliling": 0,
        "luas": 0,
        "panjang": 0,
        "lebar": 0
      }
    }
    ```
- segitiga (GET): /segitiga?alas=(number)&tinggi=(number)&hitung=(luas|keliling)
  - Response :
    ```
    {
      "message": "string",
      "data": {
        "keliling": 0,
        "luas": 0,
        "alas": 0,
        "tinggi": 0
      }
    }
    ```
- lingkaran (GET): /lingkaran?jariJari=(number)&hitung=(luas|keliling)
  - Response :
    ```
    {
      "message": "string",
      "data": {
        "keliling": 0,
        "luas": 0,
        "jariJari": 0
      }
    }
    ```

category:
- "GET" : /category
  - Response : 
    ```
    {
      "message": "string",
      "data": [
        {
          "id": 0,
          "name": "string",
          "created_date": "timestamp",
          "updated_date": "timestamp"
        },
      ]
    }
    ```
- "POST" : /category
  - Body : 
    ```
    {
      "name": "string"
    }
    ```
  - Response : 
    ```
    {
      "message": "string",
      "data": {
        "id": 0,
        "name": "string",
        "created_date": "timestamp",
        "updated_date": "timestamp"
      }
    }
    ```
- "PUT" : /category/:id
  - Body : 
    ```
    {
      "name": "string"
    }
    ```
  - Response : 
    ```
    {
      "message": "string",
      "data": {
        "id": 0,
        "name": "string",
        "created_date": "timestamp",
        "updated_date": "timestamp"
      }
    }
    ```
- "DELETE" : /category/:id
  - Response : 
    ```
    {
      "message": "string"
    }
    ```

books:
- "GET" : /category
  - Response : 
    ```
    {
      "message": "string",
      "data": [
        {
          "book_id": 0,
          "book_title": "string",
          "book_desc": "string",
          "img_url": "string",
          "book_year": 0,
          "book_price": "string",
          "book_page": 0,
          "book_thick": "string",
          "created_date": "string",
          "updated_date": "string",
          "category_id": 0
        }
      ]
    }
    ```
- "POST" : /category
  - Body : 
    ```
    {
      "category_id": 0,
      "title": "string",
      "description": "string",
      "image_url": "string",
      "release_year": 0,
      "price": "string",
      "total_page": "string"
    }
    ```
  - Response : 
    ```
    {
      "message": "string",
      "data": {
        "book_id": 0,
        "book_title": "string",
        "book_desc": "string",
        "img_url": "string",
        "book_year": 0,
        "book_price": "string",
        "book_page": 0,
        "book_thick": "string",
        "created_date": "string",
        "updated_date": "string",
        "category_id": 0
      }
    }
    ```
- "PUT" : /category/:id
  - Body : 
    ```
    {
      "category_id": 0,
      "title": "string",
      "description": "string",
      "image_url": "string",
      "release_year": 0,
      "price": "string",
      "total_page": "string"
    }
    ```
  - Response : 
    ```
    {
      "message": "string",
      "data": {
        "book_id": 0,
        "book_title": "string",
        "book_desc": "string",
        "img_url": "string",
        "book_year": 0,
        "book_price": "string",
        "book_page": 0,
        "book_thick": "string",
        "created_date": "string",
        "updated_date": "string",
        "category_id": 0
      }
    }
    ```
- "DELETE" : /category/:id
  - Response : 
    ```
    {
      "message": "string"
    }
    ```
