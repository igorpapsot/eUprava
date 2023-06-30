import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Poternica } from 'src/app/model/sudstvo/poternica';
import { PoternicaService } from 'src/app/services/sudstvo/poternica.service';

@Component({
  selector: 'app-create-poternica',
  templateUrl: './create-poternica.component.html',
  styleUrls: ['./create-poternica.component.css']
})
export class CreatePoternicaComponent implements OnInit {

  createPoternicaForm: FormGroup;
  poternicaModel: Poternica;
  sudijaId = new FormControl('');
  gradjaninId = new FormControl('');
  naslov = new FormControl('');
  opis = new FormControl('');

  constructor(private router: Router, private poternicaService: PoternicaService) {
    this.createPoternicaForm = new FormGroup({
      sudijaId: new FormControl(''),
      gradjaninId: new FormControl(''),
      naslov : new FormControl(''),
      opis : new FormControl('')
    });
    this.poternicaModel = {
      id: '',
      sudijaId: '',
      gradjaninId: '',
      naslov: '',
      opis: ''
    }
  }

  ngOnInit(){
      this.createPoternicaForm = new FormGroup({
        sudijaId: new FormControl('',Validators.required),
        gradjaninId: new FormControl('', Validators.required),
        naslov: new FormControl('',Validators.required),
        opis: new FormControl('', Validators.required)
      })
  }

  homePage(){
    this.router.navigate(['/sudstvo']);
  }

  cancel(){
    this.router.navigateByUrl('/sudstvo');
  }

  createPoternica() {
    this.poternicaModel.sudijaId = this.createPoternicaForm.get('sudijaId')?.value;
    this.poternicaModel.gradjaninId = this.createPoternicaForm.get('gradjaninId')?.value;
    this.poternicaModel.naslov = this.createPoternicaForm.get('naslov')?.value;
    this.poternicaModel.opis = this.createPoternicaForm.get('opis')?.value;
    this.poternicaService.postPoternica(this.poternicaModel).subscribe(data => {
      console.log(data);
      this.homePage();
    })
  }

}
