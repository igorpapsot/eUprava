import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Rociste } from 'src/app/model/sudstvo/rociste';
import { RocisteService } from 'src/app/services/sudstvo/rociste.service';

@Component({
  selector: 'app-create-rociste',
  templateUrl: './create-rociste.component.html',
  styleUrls: ['./create-rociste.component.css']
})
export class CreateRocisteComponent implements OnInit {

  createRocistaForm: FormGroup;
  rocisteModel: Rociste;
  datum = new FormControl('');
  mesto = new FormControl('');

  constructor(private router: Router, private rocisteService: RocisteService) {
    this.createRocistaForm = new FormGroup({
      datum: new FormControl(''),
      mesto: new FormControl('')
    });
    this.rocisteModel = {
      id: '',
      datum: '',
      mesto: '',
      sud: 2
  
    }
  }

  ngOnInit() {
      this.createRocistaForm = new FormGroup ({
        datum: new FormControl('',[Validators.required]),
        mesto: new FormControl('', Validators.required)
      })
  }

  homePage(){
    this.router.navigate(['/sudstvo']);
  }

  cancel(){
    this.router.navigateByUrl('/sudstvo');
  }

  createRociste(){
    this.rocisteModel.datum = this.createRocistaForm.get('datum')?.value;
    this.rocisteModel.mesto = this.createRocistaForm.get('mesto')?.value;
    this.rocisteService.postRociste(this.rocisteModel).subscribe(data => {
      console.log(data);
      this.homePage();
    })
  }

}
