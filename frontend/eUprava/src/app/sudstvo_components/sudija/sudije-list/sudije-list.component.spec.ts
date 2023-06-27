import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SudijeListComponent } from './sudije-list.component';

describe('SudijeListComponent', () => {
  let component: SudijeListComponent;
  let fixture: ComponentFixture<SudijeListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SudijeListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SudijeListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
