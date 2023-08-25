import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HomePageComponent } from './home-page/home-page.component';
import { HttpClientModule } from '@angular/common/http';
import { TuzilastvoComponent } from './Tuzilastvo-Komponente/tuzilastvo/tuzilastvo.component';
import { StatusPipe } from './model/tuzilastvo/statusEnum';
import { KorisnikService } from './services/mup/korisnik.service';
import { ZahtevService } from './services/mup/zahtev.service';
import { NgbdSortableHeader } from './Tuzilastvo-Komponente/tuzilastvo/tuzilastvo.component';
import { SudstvoPageComponent } from './sudstvo_components/sudstvo-page/sudstvo-page.component';

import { LogoutComponent } from './logout/logout.component';
import { MupPageComponent } from './mup-page/mup-page.component';
import { CreatePoternicaComponent } from './sudstvo_components/poternice/create-poternica/create-poternica.component';
import { ViewPoterniceComponent } from './sudstvo_components/poternice/view-poternice/view-poternice.component';
import { CreateRocisteComponent } from './sudstvo_components/rocista/create-rociste/create-rociste.component';
import { ViewRocisteComponent } from './sudstvo_components/rocista/view-rociste/view-rociste.component';
import { SideBarComponent } from './sudstvo_components/side-bar/side-bar.component';
import { CreateSudijaComponent } from './sudstvo_components/sudija/create-sudija/create-sudija.component';
import { ViewSudijeComponent } from './sudstvo_components/sudija/view-sudije/view-sudije.component';
import { ViewOptuzniceComponent } from './sudstvo_components/optuznice/view-optuznice/view-optuznice.component';
import { GranicaComponent } from './gp_components/granica/granica.component';





@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    //TuzilastvoComponent,
    StatusPipe,
    TuzilastvoComponent,
    NgbdSortableHeader,
    LogoutComponent,
    StatusPipe,
    MupPageComponent,
    SudstvoPageComponent,
    LogoutComponent,
    CreatePoternicaComponent,
    ViewPoterniceComponent,
    CreateRocisteComponent,
    ViewRocisteComponent,
    SideBarComponent,
    CreateSudijaComponent,
    ViewSudijeComponent,
    ViewOptuzniceComponent,
    GranicaComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,

  ],
  providers: [KorisnikService, ZahtevService],
  bootstrap: [AppComponent]
})
export class AppModule { }
