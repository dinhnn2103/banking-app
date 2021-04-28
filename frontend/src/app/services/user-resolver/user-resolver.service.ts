import {ActivatedRouteSnapshot, Resolve, RouterStateSnapshot} from '@angular/router';
import {UserService} from '../user/user.service';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserResolverService implements Resolve<any> {

  constructor(private user: UserService) {
  }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any | Observable<never> {
    return this.user.getUser();
  }

}
