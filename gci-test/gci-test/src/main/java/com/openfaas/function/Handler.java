package com.openfaas.function;

import com.openfaas.model.IHandler;
import com.openfaas.model.IResponse;
import com.openfaas.model.IRequest;
import com.openfaas.model.Response;

import java.util.List;
import java.util.ArrayList;

public class Handler implements com.openfaas.model.IHandler {

    public IResponse Handle(IRequest req) {
        Response res = new Response();
        int size = (int) Math.pow(2, 21);
        List<Integer> list  = new ArrayList();
        for (int i = 0; i < size; i++) {
            list.add(i);
        }
        
        res.setBody("GCI-Test Invoked!");
	    return res;
    }
}
