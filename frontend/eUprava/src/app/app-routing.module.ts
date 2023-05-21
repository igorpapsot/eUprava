import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { TuzilastvoComponent } from './Tuzilastvo-Komponente/tuzilastvo/tuzilastvo.component';
import { MupPageComponent } from './mup-page/mup-page.component';

const routes: Routes = [

  {path: '', component: HomePageComponent},
  {path: 'tuzilastvo', component: TuzilastvoComponent},
  {path: 'mup', component: MupPageComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes, {onSameUrlNavigation:"reload"})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
