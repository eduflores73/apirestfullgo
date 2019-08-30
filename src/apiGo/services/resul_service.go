package services

import (
	"../domains"
	"../utils"
	"sync"
)

func GetResult(userId int) (*domains.Result, *utils.ApiError){

	user := &domains.User{
		ID:userId,
	}
	if err := user.Get(); err != nil{
		return nil, err
	}

	country := &domains.Country{
		ID: user.CountryID,
	}
	site := &domains.Site{
		ID:	user.SiteID,
	}
	if err := country.Get(); err != nil{
		return nil, err

	}
	if err := site.Get(); err != nil{
		return nil, err
	}

	resp := &domains.Result{
		User : user,
		Site : site,
		Country: country,
	}

	return resp, nil
}

func GetResultWG(userId int) (*domains.Result, *utils.ApiError){

	apiError := utils.ApiError{}
	user := &domains.User{
		ID:userId,
	}
	user.Get()

	country := &domains.Country{
		ID: user.CountryID,
	}
	site := &domains.Site{
		ID:	user.SiteID,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go country.GetWg(&wg,&apiError)
	go site.GetWG(&wg,&apiError)

	wg.Wait()

	if apiError.Status != 0{
		return nil, &apiError
	}

	resp := &domains.Result{
		User : user,
		Site : site,
		Country: country,
	}

	return resp, nil
}

func GetResultCH(userId int) (*domains.Result, *utils.ApiError){

	user := &domains.User{
		ID:userId,
	}
	user.Get()

	country := &domains.Country{
		ID: user.CountryID,
	}
	site := &domains.Site{
		ID:	user.SiteID,
	}

	valorCH := make(chan domains.Result, 2)
	resultadoFinal := &domains.Result{
		User:user,
	}
	go country.GetCH(valorCH)
	go site.GetCH(valorCH)
	for i :=0; i < 2; i++{
		result := <- valorCH
		if result.Error != nil{
			return nil,result.Error
		}
		if result.Country != nil{
			resultadoFinal.Country = result.Country
		}
		if result.Site != nil{
			resultadoFinal.Site = result.Site
		}
	}

	return resultadoFinal, nil
}