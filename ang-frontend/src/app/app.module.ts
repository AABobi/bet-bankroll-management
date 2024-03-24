import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginPageComponent } from './home-page/login-page/login-page.component';
import { RegisterPageComponent } from './home-page/register-page/register-page.component';
import { HomeViewComponent } from './home-page/home-view/home-view.component';
import { StoreModule } from '@ngrx/store';
import { simpleReducer } from './simple.reducer';

@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    RegisterPageComponent,
    HomeViewComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    StoreModule.forRoot({ message: simpleReducer }, {}),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
