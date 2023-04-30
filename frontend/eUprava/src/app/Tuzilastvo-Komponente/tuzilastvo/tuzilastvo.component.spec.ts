import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TuzilastvoComponent } from './tuzilastvo.component';

describe('TuzilastvoComponent', () => {
  let component: TuzilastvoComponent;
  let fixture: ComponentFixture<TuzilastvoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TuzilastvoComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TuzilastvoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
