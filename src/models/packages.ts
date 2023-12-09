export namespace PackageModel {
  export class Package {
    name: string;
    version: string;
    owner: string;
    description: string;

    constructor(
      name: string,
      version: string,
      desctription: string,
      owner: string
    ) {
      this.name = name;
      this.version = version;
      this.description = desctription;
      this.owner = owner;
    }
  }
  export async function search(query: string, useDescription: boolean = true) {
    const packageDB = (await import("./db/packages.json")).default;
    const result: Package[] = [];
    (Object.keys(packageDB) as (keyof typeof packageDB)[]).forEach(
      (pkgName) => {
        const pkg = new Package(
          pkgName,
          packageDB[pkgName].version,
          packageDB[pkgName].description,
          packageDB[pkgName].owner
        );
        if (
          pkgName.toLowerCase().includes(query.toLowerCase()) ||
          (useDescription &&
            pkg.description.toLowerCase().includes(query.toLowerCase()))
        ) {
          result.push(pkg);
        }
      }
    );
    return result;
  }
}
