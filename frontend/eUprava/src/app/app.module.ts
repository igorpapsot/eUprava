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
import { PoternicaComponent } from './sudstvo_components/poternica/poternica.component';
import { RocisteComponent } from './sudstvo_components/rociste/rociste.component';
@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    //TuzilastvoComponent,
    StatusPipe,
    TuzilastvoComponent,
    NgbdSortableHeader,
    SudstvoPageComponent,
    LogoutComponent,
    PoternicaComponent,
    RocisteComponent
    
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
