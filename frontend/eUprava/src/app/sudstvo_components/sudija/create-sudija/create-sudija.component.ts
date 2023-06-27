import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Sudija } from 'src/app/model/sudstvo/sudija';
import { SudijaService } from 'src/app/services/sudstvo/sudija.service';

@Component({
  selector: 'app-create-sudija',
  templateUrl: './create-sudija.component.html',
  styleUrls: ['./create-sudija.component.css']
})
export class CreateSudijaComponent implements OnInit {

  createSudijaForm: FormGroup;
  sudijaModel: Sudija;
  ime = new FormControl('');
  prezime = new FormControl('');
  jmbg = new FormControl('');
  lozinka = new FormControl('');

  constructor(private router: Router, private sudijaService: SudijaService) {
    this.createSudijaForm = new FormGroup({
      ime: new FormControl(''),
      prezime: new FormControl(''),
      jmbg: new FormControl(''),
      lozinka: new FormControl('')
    });

    this.sudijaModel = {
      id: '',
      ime: '',
      prezime: '',
      pol: 1,
      jmbg: '',
      lozinka: '',
      sud: 2
    }
  }

  ngOnInit() {
      this.createSudijaForm = new FormGroup({
        ime: new FormControl('', Validators.required),
        prezime: new FormControl('', Validators.required),
        jmbg: new FormControl('', Validators.required),
        lozinka: new FormControl('', Validators.required)

      })
  }

  homePage(){
    this.router.navigate(['/sudstvo']);
  }

  cancel(){
    this.router.navigateByUrl('/sudstvo');
  }

  createSudija() {
    this.sudijaModel.ime = this.createSudijaForm.get('ime')?.value;
    this.sudijaModel.prezime = this.createSudijaForm.get('prezime')?.value;
    this.sudijaModel.jmbg = this.createSudijaForm.get('jmbg')?.value;
    this.sudijaModel.lozinka = this.createSudijaForm.get('lozinka')?.value;
    this.sudijaService.registerSudija(this.sudijaModel).subscribe(data => {
      console.log(data);
      this.homePage();
    })
  }

  

}
