- StockPro Project Description

Follow these setup to run the Project:

    1. Get the code from Github: 
                $cd go/src
                $ git clone https://github.com/Nagarjunhabbu/StockPro.git
            
    2. Run the Server: $go run main.go
  
  -------------------------------------------------------------------------------------------------------------------

  API Information-

           To perform the operation use below link


    1. GET request- http://localhost:8000/api/v1/investment   -  (/investment - To get the Top 5 best investments irrespective of invType)

    2. GET request- http://localhost:8000/api/v1/investment?invType="Stocks"   -  (pass invType-optional Param - To get the Top 5 best investments in that particular invType)

    3. POST Request - http://localhost:8000/api/v1/investment - (To get portfolio suggestion send JSON body along with this { "ERoI":7,"Amount":10000 })- This will return the portfolio suggestion for investment that satisfies give RoI irrespective of the type

    4. POST Request - http://localhost:8000/api/v1/investment - (To get portfolio suggestion in specified investment type send JSON body along with this { "invType":"Stocks","ERoI":7,"Amount":10000 })- This will return the portfolio suggestion for investment that satisfies give RoI in that particular Investment type


   