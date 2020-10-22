import React, { useState, useEffect } from 'react';
import {
  pageTheme,
  Page,
  Header,
  Content,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { EntMeal } from '../../api/models/EntMeal';
import { EntFood } from '../../api/models/EntFood';
import { EntUser } from '../../api/models/EntUser';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    textField: {
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 200,
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 200,
    },
    button: {
      margin: theme.spacing(1),
      textAlign: 'center',
    },
  }),
);

export default function MealPlan() {

  const classes = useStyles();

  const api = new DefaultApi();

  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [meals, setMeals] = React.useState<EntMeal[]>([]);
  const [foods, setFoods] = React.useState<EntFood[]>([]);

  const [date, setDate] = React.useState(String);
  const [userid, setUser] = React.useState(Number);
  const [foodid, setFood] = React.useState(Number);
  const [mealid, setMeal] = React.useState(Number);

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);
  const [loading, setLoading] = React.useState(true);

  const getUsers = async () => {
    const res = await api.listUser({ limit: 10, offset: 0 });
    setUsers(res);
    setLoading(false);
    console.log(res);
  };

  const getMeals = async () => {
    const res = await api.listMeal({ limit: 10, offset: 0 });
    setMeals(res);  
    setLoading(false);
    console.log(res);
  };
  const getFoods = async () => {
    const res = await api.listFood({ limit: 10, offset: 0 });
    setFoods(res);  
    setLoading(false);
    console.log(res);
  };

  useEffect(() => {
    getUsers();
    getMeals();
    getFoods();
  }, [loading]);

  const CreateMealplan = async () => {
    const mealplan = {
      date: date + "T00:00:00+07:00",
      food: foodid,
      meal: mealid,
      owner: userid,
    };
    console.log(mealplan);
    const res: any = await api.createMealplan({ mealplan : mealplan });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      window.location.reload(false);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const handleDateChange = (event: any) => {
    setDate(event.target.value as string);
  };

  const handleUserChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const handleFoodChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFood(event.target.value as number);
  };

  const handleMealChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setMeal(event.target.value as number);
  };

  return(
    <Page theme={pageTheme.home}>
      <Header title={`MEAL PLANNER`}></Header>
      <Content>
        <ContentHeader title="Create Plan">
        {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 Success!!
               </Alert>
              ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 Error!!
               </Alert>
             )}
           </div>
         ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Paper className={classes.paper}>
                <div>
                  <FormControl className={classes.formControl}>
                    <InputLabel id="demo-simple-select-label">เลือกผู้ใช้ระบบ</InputLabel>
                    <Select
                      name="owner"
                      value={userid}
                      onChange={handleUserChange}
                    >
                    {users.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.email}
                        </MenuItem>
                      );
                    })}
                    </Select>
                  </FormControl>
                </div>
                <Grid item xs={12}>
                  <div className={classes.paper}></div>
                </Grid>
                <div>
                <TextField
                  label="เลือกวันที่"
                  name="date"
                  type="date"
                  value={date}
                  className={classes.textField}
                  onChange={handleDateChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
                </div>
                <Grid item xs={12}>
                  <div className={classes.paper}></div>
                </Grid>
                <div>
                <FormControl className={classes.formControl}>
                  <InputLabel id="demo-simple-select-label">เลือกมื้ออาหาร</InputLabel>
                  <Select
                    name="meal"
                    value={mealid}
                    onChange={handleMealChange}
                  >
                  {meals.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                  </Select>
                </FormControl>
                </div>
                <Grid item xs={12}>
                  <div className={classes.paper}></div>
                </Grid>
                <div>
                <FormControl className={classes.formControl}>
                  <InputLabel id="demo-simple-select-label">เลือกรายการอาหาร</InputLabel>
                  <Select
                    name="food"
                    value={foodid}
                    onChange={handleFoodChange}
                  >
                  {foods.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                  </Select>
                </FormControl>
                </div>
                <Grid item xs={12}>
                  <div className={classes.paper}></div>
                </Grid>
                <div>
                  <Button
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    onClick={() => { 
                      CreateMealplan(); 
                    }}
                    startIcon={<AddIcon />}
                  >
                    เพิ่มรายการอาหาร
                  </Button>
                </div>
              </Paper>
            </Grid>
          </Grid>
        </div>
      </Content>
    </Page>
  );
}
