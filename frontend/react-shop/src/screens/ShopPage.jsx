import React, {useEffect, useState} from 'react';
import {getProducts} from "./requests/getProducts";
import {CardMedia, CardActions, Button, Grid, Card, CardContent, Typography} from "@mui/material";
import ShoppingBasketIcon from '@mui/icons-material/ShoppingBasket';
import { useCart, useDispatchCart } from "./components/Card";

const ShopPage = () => {
    const dispatch = useDispatchCart();
    const [products, setProducts] = useState([])
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");


    useEffect(() => {
        getProducts(setLoading, setProducts, setError)
    },[]);



    const handleAddToCard = (id) => {
        console.log(id)
    }

    const addToCart = (item) => {
        dispatch({ type: "ADD", item });
    };



    return(
        <div style={{padding: "50px"}}>{
            loading
                ?
                <div>Loading ...</div>
                :
                <Grid container spacing={2} >
                    {
                        products.map((prod, index) => {
                            return (
                           <Grid item xs={4}>
                                   <Card sx={{ maxWidth: 350}}>
                                      <CardMedia
                                        component="img"
                                        height="200"
                                        image="https://images.unsplash.com/photo-1602573852058-ef7c665fcd92?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=774&q=80"
                                        alt="green iguana"
                                      />
                                      <CardContent>
                                        <Typography gutterBottom variant="h5" component="div">
                                            {prod.Categorie}
                                        </Typography>
                                        <Typography variant="body2" color="text.secondary">
                                             {prod.Name}
                                        </Typography>
                                      </CardContent>
                                      <CardActions>
                                        <ShoppingBasketIcon style={{color: "black", padding: "4px"}} />
                                        <Button sx={{ my: 2, color: "black" }} size="small" onClick={() => addToCart(prod)}>Add to Card</Button>
                                      </CardActions>
                                    </Card>
                           </Grid>
                            )
                        })
                    }
                </Grid>

        }
        </div>
    )
}

export default ShopPage;